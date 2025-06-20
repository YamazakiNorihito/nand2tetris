namespace BasicAssembler.Tests;

public class CodeTests
{
    private readonly ICode _code = new Code();

    [Theory]
    [InlineData("", "000")]
    [InlineData("M", "001")]
    [InlineData("D", "010")]
    [InlineData("MD", "011")]
    [InlineData("A", "100")]
    [InlineData("AM", "101")]
    [InlineData("AD", "110")]
    [InlineData("AMD", "111")]
    public void Dest_ReturnsExpectedBinary(string input, string expected)
    {
        Assert.Equal(expected, _code.Dest(input));
    }

    [Theory]
    [InlineData("0", "0101010")]
    [InlineData("1", "0111111")]
    [InlineData("-1", "0111010")]
    [InlineData("D", "0001100")]
    [InlineData("A", "0110000")]
    [InlineData("!D", "0001101")]
    [InlineData("!A", "0110001")]
    [InlineData("-D", "0001111")]
    [InlineData("-A", "0110011")]
    [InlineData("D+1", "0011111")]
    [InlineData("A+1", "0110111")]
    [InlineData("D-1", "0001110")]
    [InlineData("A-1", "0110010")]
    [InlineData("D+A", "0000010")]
    [InlineData("D-A", "0010011")]
    [InlineData("A-D", "0000111")]
    [InlineData("D&A", "0000000")]
    [InlineData("D|A", "0010101")]
    [InlineData("M", "1110000")]
    [InlineData("!M", "1110001")]
    [InlineData("-M", "1110011")]
    [InlineData("M+1", "1110111")]
    [InlineData("M-1", "1110010")]
    [InlineData("D+M", "1000010")]
    [InlineData("D-M", "1010011")]
    [InlineData("M-D", "1000111")]
    [InlineData("D&M", "1000000")]
    [InlineData("D|M", "1010101")]
    public void Comp_ReturnsExpectedBinary(string input, string expected)
    {
        Assert.Equal(expected, _code.Comp(input));
    }

    [Theory]
    [InlineData("", "000")]
    [InlineData("JGT", "001")]
    [InlineData("JEQ", "010")]
    [InlineData("JGE", "011")]
    [InlineData("JLT", "100")]
    [InlineData("JNE", "101")]
    [InlineData("JLE", "110")]
    [InlineData("JMP", "111")]
    public void Jump_ReturnsExpectedBinary(string input, string expected)
    {
        Assert.Equal(expected, _code.Jump(input));
    }
}