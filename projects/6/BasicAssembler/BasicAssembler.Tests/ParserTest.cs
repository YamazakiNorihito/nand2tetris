namespace BasicAssembler.Tests;

public class ParserTests
{

    [Fact]
    public void CreateParser_WithAbsolutePath_ReturnsParser()
    {
        using var temp = new TempFile(["@2", "D=A"]);
        var parser = Parser.CreateParser(temp.Path);
        Assert.NotNull(parser);
    }

    [Fact]
    public void CreateParser_WithRelativePath_ReturnsParser()
    {
        var fileName = "testRelative.asm";
        using var temp = new TempFile(["@5"], fileName);
        var parser = Parser.CreateParser(fileName);
        Assert.NotNull(parser);
    }

    [Fact]
    public void HasMoreLines_WithOnlyEmptyLines_ReturnsFalse()
    {
        var lines = Array.Empty<string>();
        using var temp = new TempFile(lines);
        var parser = Parser.CreateParser(temp.Path);

        Assert.False(parser.HasMoreLines());
    }

    [Fact]
    public void HasMoreLines_WithValidInstructionLines_ReturnsTrue()
    {
        var lines = new[]
        {
            "@10",
            "D=A"
        };
        using var temp = new TempFile(lines);
        var parser = Parser.CreateParser(temp.Path);

        Assert.True(parser.HasMoreLines());
    }

    [Theory]
    [InlineData(new string[] { "// comment", "", "   ", "@1", "D=M" }, "@1")]
    [InlineData(new string[] { "@2" }, "@2")]
    [InlineData(new string[] { "   // only comment", "   ", "" }, "")]
    [InlineData(new string[] { "   // only comment1", "   // only comment2" }, "")]
    [InlineData(new string[] { "   // only comment1", "@3   // only comment2" }, "@3")]
    public void Advance_SetsCurrentLineToSecondInstruction(string[] inputLines, string expected)
    {
        using var temp = new TempFile(inputLines);

        var parser = Parser.CreateParser(temp.Path);
        parser.Advance();

        Assert.Equal(expected, parser.CurrentLine());
    }

    [Theory]
    [InlineData(new string[] { "// comment", "", "   ", "@1", "D=M" }, "@1", "D=M")]
    [InlineData(new string[] { "@2", "@3" }, "@2", "@3")]
    [InlineData(new string[] { "   // only comment", "   ", "@5" }, "@5", "")]
    [InlineData(new string[] { "@7   // comment", "M=D", "// end" }, "@7", "M=D")]
    public void Advance_Twice_SetsCurrentLineToSecondInstruction(string[] inputLines, string expectedFirst, string expectedSecond)
    {
        using var temp = new TempFile(inputLines);

        var parser = Parser.CreateParser(temp.Path);

        parser.Advance();
        Assert.Equal(expectedFirst, parser.CurrentLine());

        parser.Advance();
        Assert.Equal(expectedSecond, parser.CurrentLine());
    }

    [Theory]
    [InlineData(new[] { "@2" }, "A")]
    [InlineData(new[] { "   @15   // set value" }, "A")]
    [InlineData(new[] { "(LOOP)" }, "L")]
    [InlineData(new[] { "   (END)   " }, "L")]
    [InlineData(new[] { "D=M" }, "C")]
    [InlineData(new[] { "M=D;JGT" }, "C")]
    [InlineData(new[] { "// pure comment", "", "   " }, "")]
    [InlineData(new[] { "// comment", "@3 // inline comment" }, "A")]
    public void InstructionType_ReturnsCorrectInstructionType(string[] inputLines, string expectedType)
    {
        using var temp = new TempFile(inputLines);
        var parser = Parser.CreateParser(temp.Path);

        parser.Advance();
        Assert.Equal(expectedType, parser.InstructionType());
    }

    [Theory]
    [InlineData(new[] { "@21" }, "21")]
    [InlineData(new[] { "@LOOP" }, "LOOP")]
    [InlineData(new[] { "   @R2   // comment" }, "R2")]
    [InlineData(new[] { "(END)" }, "END")]
    [InlineData(new[] { "   (START)   " }, "START")]
    [InlineData(new[] { "// comment", "@foo // inline" }, "foo")]
    [InlineData(new[] { "D=M" }, "")]
    public void Symbol_ReturnsCorrectSymbol(string[] inputLines, string expectedSymbol)
    {
        using var temp = new TempFile(inputLines);
        var parser = Parser.CreateParser(temp.Path);

        parser.Advance();
        Assert.Equal(expectedSymbol, parser.Symbol());
    }

    [Theory]
    [InlineData(new[] { "D=M" }, "D")]
    [InlineData(new[] { "M=D" }, "M")]
    [InlineData(new[] { "MD=D+1" }, "MD")]
    [InlineData(new[] { "A=M-1" }, "A")]
    [InlineData(new[] { "D;JGT" }, "")]
    [InlineData(new[] { "D+1" }, "")]
    [InlineData(new[] { "@2" }, "")]
    [InlineData(new[] { "(LOOP)" }, "")]
    public void Dest_ReturnsCorrectDest(string[] inputLines, string expectedDest)
    {
        using var temp = new TempFile(inputLines);
        var parser = Parser.CreateParser(temp.Path);

        parser.Advance();
        Assert.Equal(expectedDest, parser.Dest());
    }

    [Theory]
    [InlineData(new[] { "D=M" }, "M")]
    [InlineData(new[] { "M=D" }, "D")]
    [InlineData(new[] { "MD=D+1" }, "D+1")]
    [InlineData(new[] { "A=M-1" }, "M-1")]
    [InlineData(new[] { "D+1" }, "D+1")]
    [InlineData(new[] { "D;JGT" }, "D")]
    [InlineData(new[] { "0;JMP" }, "0")]
    [InlineData(new[] { "@2" }, "")]
    [InlineData(new[] { "(LOOP)" }, "")]
    public void Comp_ReturnsCorrectComp(string[] inputLines, string expectedComp)
    {
        using var temp = new TempFile(inputLines);
        var parser = Parser.CreateParser(temp.Path);

        parser.Advance();
        Assert.Equal(expectedComp, parser.Comp());
    }

    [Theory]
    [InlineData(new[] { "D=M" }, "")]
    [InlineData(new[] { "M=D" }, "")]
    [InlineData(new[] { "MD=D+1" }, "")]
    [InlineData(new[] { "A=M-1" }, "")]
    [InlineData(new[] { "D+1" }, "")]
    [InlineData(new[] { "D;JGT" }, "JGT")]
    [InlineData(new[] { "0;JMP" }, "JMP")]
    [InlineData(new[] { "@2" }, "")]
    [InlineData(new[] { "(LOOP)" }, "")]
    public void Jump_ReturnsCorrectJump(string[] inputLines, string expectedJump)
    {
        using var temp = new TempFile(inputLines);
        var parser = Parser.CreateParser(temp.Path);

        parser.Advance();
        Assert.Equal(expectedJump, parser.Jump());
    }
}