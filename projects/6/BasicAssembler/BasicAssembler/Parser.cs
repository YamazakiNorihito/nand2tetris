using System.Text.RegularExpressions;

namespace BasicAssembler;

public interface IParser
{
    public const string A_INSTRUCTION = "A";
    public const string C_INSTRUCTION = "C";
    public const string L_INSTRUCTION = "L";

    string CurrentLine();

    bool HasMoreLines();

    void Advance();

    string InstructionType();

    string Symbol();

    string Dest();

    string Comp();

    string Jump();
}

public class Parser : IParser
{
    private string[] _lines;
    private int _currentLineIndex;
    private string _currentLine;


    private string _instructionType;
    private string _symbol;
    private string _dest;
    private string _comp;
    private string _jump;

    private static readonly Regex _aCommandRegex =
        new(pattern: @"^@([A-Za-z_.$:][A-Za-z0-9_.$:]*|\d+)$", options: RegexOptions.Compiled);

    private static readonly Regex _lCommandRegex =
        new(pattern: @"^\(([A-Za-z_.$:][A-Za-z0-9_.$:]*)\)$", options: RegexOptions.Compiled);

    // C 命令: [dest=]comp[;jump]
    private static readonly Regex _cCommandRegex = new(
        pattern: @"
            ^
            (?:(?<dest>[AMD]{1,3})=)?
            (?<comp>
                0|1|-1
                |D|A|M
                |!D|!A|!M
                |-D|-A|-M
                |D\+1|A\+1|M\+1
                |D-1|A-1|M-1
                |D\+A|D\+M
                |D-A|D-M
                |A-D|M-D
                |D&A|D&M
                |D\|A|D\|M
            )
            (?:;(?<jump>JGT|JEQ|JGE|JLT|JNE|JLE|JMP))?
            $
        ",
        options: RegexOptions.Compiled | RegexOptions.IgnorePatternWhitespace
    );


    protected Parser(string[] assemblyLines)
    {
        _lines = assemblyLines;
        _currentLineIndex = 0;
        _currentLine = string.Empty;
        _instructionType = string.Empty;
        _symbol = string.Empty;
        _dest = string.Empty;
        _comp = string.Empty;
        _jump = string.Empty;
    }

    public static IParser CreateParser(string assemblyFilePath)
    {
        string path = assemblyFilePath;
        if (!Path.IsPathRooted(assemblyFilePath))
        {
            path = Path.Combine(Directory.GetCurrentDirectory(), assemblyFilePath);
        }
        var lines = File.ReadAllLines(path);
        return new Parser(lines);
    }

    public string CurrentLine() => _currentLine;

    public bool HasMoreLines()
    {
        return _currentLineIndex < _lines.Length;
    }

    public void Advance()
    {
        _currentLine = string.Empty;
        _instructionType = string.Empty;
        _symbol = string.Empty;
        _dest = string.Empty;
        _comp = string.Empty;
        _jump = string.Empty;

        while (HasMoreLines())
        {
            var line = _lines[_currentLineIndex++].Trim();

            // コメント行や空行をまとめてスキップ
            if (line.Length == 0 || line.StartsWith("//"))
                continue;

            int commentIndex = line.IndexOf("//");
            if (commentIndex >= 0)
                line = line[..commentIndex].Trim();

            if (line.Length == 0)
                continue;

            _currentLine = line;
            ParseInstruction(line);
            break;
        }
    }

    public string InstructionType() => _instructionType;

    public string Symbol() => _symbol;

    public string Dest() => _dest;

    public string Comp() => _comp;

    public string Jump() => _jump;

    private void ParseInstruction(string line)
    {
        var aCommandMatch = _aCommandRegex.Match(line);
        if (aCommandMatch.Success)
        {
            _instructionType = IParser.A_INSTRUCTION;
            _symbol = aCommandMatch.Groups[1].Value;
            return;
        }

        var lCommandMatch = _lCommandRegex.Match(line);
        if (lCommandMatch.Success)
        {
            _instructionType = IParser.L_INSTRUCTION;
            _symbol = lCommandMatch.Groups[1].Value;
            return;
        }

        var cCommandMatch = _cCommandRegex.Match(line);
        if (cCommandMatch.Success)
        {
            _instructionType = IParser.C_INSTRUCTION;
            _dest = cCommandMatch.Groups["dest"].Value;
            _comp = cCommandMatch.Groups["comp"].Value;
            _jump = cCommandMatch.Groups["jump"].Value;
            return;
        }

        throw new InvalidOperationException($"Invalid instruction format: {line}");
    }
}