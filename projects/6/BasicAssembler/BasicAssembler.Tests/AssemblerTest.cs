using System.Text;
using Moq;

namespace BasicAssembler.Tests;

public class AssemblerTest
{
    [Fact]
    public void Assemble_WritesAInstructionBinary()
    {
        // Arrange
        var parser = new Mock<IParser>();
        parser.SetupSequence(p => p.HasMoreLines())
            .Returns(true)
            .Returns(false);
        parser.Setup(p => p.Advance());
        parser.Setup(p => p.InstructionType()).Returns(IParser.A_INSTRUCTION);
        parser.Setup(p => p.Symbol()).Returns("2");

        var assembler = new Assembler(Mock.Of<ICode>(), parser.Object);

        using var ms = new MemoryStream();
        using var writer = new StreamWriter(ms, Encoding.UTF8, 1024, true);

        // Act
        assembler.Assemble(writer);
        writer.Flush();
        ms.Position = 0;
        var output = new StreamReader(ms).ReadToEnd();

        // Assert
        // 16-bit A-instruction for @2
        Assert.Contains("0000000000000010", output);
    }

    [Fact]
    public void Assemble_WritesCInstructionBinary()
    {
        // Arrange
        var parser = new Mock<IParser>();
        parser.SetupSequence(p => p.HasMoreLines())
            .Returns(true)
            .Returns(false);
        parser.Setup(p => p.Advance());
        parser.Setup(p => p.InstructionType()).Returns(IParser.C_INSTRUCTION);
        parser.Setup(p => p.Dest()).Returns("D");
        parser.Setup(p => p.Comp()).Returns("A");
        parser.Setup(p => p.Jump()).Returns("JGT");

        var code = new Mock<ICode>();
        code.Setup(c => c.Dest("D")).Returns("010");
        code.Setup(c => c.Comp("A")).Returns("0110000");
        code.Setup(c => c.Jump("JGT")).Returns("001");

        var assembler = new Assembler(code.Object, parser.Object);

        using var ms = new MemoryStream();
        using var writer = new StreamWriter(ms, Encoding.UTF8, 1024, true);

        // Act
        assembler.Assemble(writer);
        writer.Flush();
        ms.Position = 0;
        var output = new StreamReader(ms).ReadToEnd();

        // Assert
        // 16-bit C-instruction for D=A;JGT
        Assert.Contains("1110110000010001", output);
    }

    [Fact]
    public void Assemble_DoesNotWriteForLInstruction()
    {
        // Arrange
        var parser = new Mock<IParser>();
        parser.SetupSequence(p => p.HasMoreLines())
            .Returns(true)
            .Returns(false);
        parser.Setup(p => p.Advance());
        parser.Setup(p => p.InstructionType()).Returns(IParser.L_INSTRUCTION);
        parser.Setup(p => p.Symbol()).Returns("LOOP");

        var assembler = new Assembler(Mock.Of<ICode>(), parser.Object);

        using var ms = new MemoryStream();
        using var writer = new StreamWriter(ms, Encoding.UTF8, 1024, true);

        // Act
        assembler.Assemble(writer);
        writer.Flush();
        ms.Position = 0;
        var output = new StreamReader(ms).ReadToEnd();

        // Assert
        // Synbol is not written to output
        Assert.True(string.IsNullOrEmpty(output));
    }


    /*
        @2
        D=A
        @3
        D=D+A
        @0
        M=D
    */
    [Fact]
    public void Assemble_ComputesR0Equals2Plus3()
    {
        // Arrange
        var parser = new Mock<IParser>();
        parser.SetupSequence(p => p.HasMoreLines())
            .Returns(true).Returns(true).Returns(true).Returns(true).Returns(true).Returns(true).Returns(false);
        parser.SetupSequence(p => p.InstructionType())
            .Returns(IParser.A_INSTRUCTION) // @2
            .Returns(IParser.C_INSTRUCTION) // D=A
            .Returns(IParser.A_INSTRUCTION) // @3
            .Returns(IParser.C_INSTRUCTION) // D=D+A
            .Returns(IParser.A_INSTRUCTION) // @0
            .Returns(IParser.C_INSTRUCTION); // M=D
        parser.SetupSequence(p => p.Symbol())
            .Returns("2")
            .Returns("3")
            .Returns("0");
        parser.SetupSequence(p => p.Dest())
            .Returns("D")
            .Returns("D")
            .Returns("M");
        parser.SetupSequence(p => p.Comp())
            .Returns("A")
            .Returns("D+A")
            .Returns("D");
        parser.Setup(c => c.Jump()).Returns("");

        parser.Setup(p => p.Advance());

        var code = new Mock<ICode>();
        code.Setup(c => c.Comp("A")).Returns("0110000");
        code.Setup(c => c.Comp("D+A")).Returns("0000010");
        code.Setup(c => c.Comp("D")).Returns("0001100");
        code.Setup(c => c.Dest("D")).Returns("010");
        code.Setup(c => c.Dest("M")).Returns("001");
        code.Setup(c => c.Jump(It.IsAny<string>())).Returns("000");

        var assembler = new Assembler(code.Object, parser.Object);

        using var ms = new MemoryStream();
        using var writer = new StreamWriter(ms, Encoding.UTF8, 1024, true);

        // Act
        assembler.Assemble(writer);
        writer.Flush();
        ms.Position = 0;
        var output = new StreamReader(ms).ReadToEnd();

        // Assert: 全出力をまとめて比較
        var expected = string.Join('\n', new[]
        {
        "0000000000000010",
        "1110110000010000",
        "0000000000000011",
        "1110000010010000",
        "0000000000000000",
        "1110001100001000"
    }) + "\n";

        Assert.Equal(expected, output);
    }
}