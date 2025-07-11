public interface IParser
{
    bool HasMoreLines();
    bool Advance();
    CommandType CommandType { get; }
    string Arg1 { get; }
    int Arg2 { get; }
    int RawLineLength();
    int CurrentLineIndex();
}

public class Parser : IParser
{
    private readonly List<string> lines;
    private int currentLineIndex;
    private string currentLine = "";
    private int rawLineLength;

    private CommandType commandType;
    private string arg1 = "";
    private int arg2;

    public Parser(TextReader reader)
    {
        var content = reader.ReadToEnd();
        if (string.IsNullOrEmpty(content))
        {
            lines = new List<string>();
            currentLineIndex = -1;
            commandType = CommandType.C_UNKNOWN;
            return;
        }

        var rawLines = content.Replace("\r\n", "\n").Split('\n');
        rawLineLength = rawLines.Length;

        lines = new List<string>();
        foreach (var line in rawLines)
        {
            var trimmed = line;
            var commentIndex = trimmed.IndexOf("//", StringComparison.Ordinal);
            if (commentIndex != -1)
            {
                trimmed = trimmed.Substring(0, commentIndex);
            }
            trimmed = trimmed.Trim();
            if (!string.IsNullOrEmpty(trimmed))
            {
                lines.Add(trimmed);
            }
        }
        commandType = CommandType.C_UNKNOWN;
        currentLineIndex = -1;
    }

    public bool HasMoreLines()
    {
        return currentLineIndex + 1 < lines.Count;
    }

    public bool Advance()
    {
        if (!HasMoreLines())
        {
            return false;
        }
        ResetInstructionParts();
        currentLineIndex++;
        currentLine = lines[currentLineIndex];

        var parts = currentLine.Split([' ', '\t'], StringSplitOptions.RemoveEmptyEntries);
        if (parts.Length == 0)
        {
            throw new Exception("empty command line");
        }

        var command = parts[0];

        var commandInfo = CommandConfig.GetCommandInfo(command);
        if (commandInfo.type == CommandType.C_UNKNOWN)
        {
            throw new Exception($"unknown command: {command}");
        }

        if (parts.Length - 1 != commandInfo.argLength)
        {
            throw new Exception($"invalid argument count for {command}: {currentLine}");
        }

        commandType = commandInfo.type;
        arg1 = commandInfo.argLength >= 1 ? parts[1] : command;
        arg2 = commandInfo.argLength == 2 ? int.Parse(parts[2]) : 0;
        return true;
    }

    public CommandType CommandType => commandType;
    public string Arg1 => arg1;
    public int Arg2 => arg2;

    private void ResetInstructionParts()
    {
        currentLine = "";
        commandType = CommandType.C_UNKNOWN;
        arg1 = "";
        arg2 = 0;
    }

    public int RawLineLength() => rawLineLength;
    public int CurrentLineIndex() => currentLineIndex;
}
