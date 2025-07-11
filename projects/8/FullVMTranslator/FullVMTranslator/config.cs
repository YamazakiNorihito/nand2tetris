public static class CommandConfig
{
    private static readonly Dictionary<string, (CommandType type, int argLength)> commandInfo = new()
    {
            {"push", (CommandType.C_PUSH, 2)},
            {"pop", (CommandType.C_POP, 2)},
            {"label", (CommandType.C_LABEL, 1)},
            {"goto", (CommandType.C_GOTO, 1)},
            {"if-goto", (CommandType.C_IF, 1)},
            {"function", (CommandType.C_FUNCTION, 2)},
            {"call", (CommandType.C_CALL, 2)},
            {"return", (CommandType.C_RETURN, 0)},
            {"add", (CommandType.C_ARITHMETIC, 0)},
            {"sub", (CommandType.C_ARITHMETIC, 0)},
            {"neg", (CommandType.C_ARITHMETIC, 0)},
            {"and", (CommandType.C_ARITHMETIC, 0)},
            {"or", (CommandType.C_ARITHMETIC, 0)},
            {"not", (CommandType.C_ARITHMETIC, 0)},
            {"eq", (CommandType.C_ARITHMETIC, 0)},
            {"gt", (CommandType.C_ARITHMETIC, 0)},
            {"lt", (CommandType.C_ARITHMETIC, 0)}
        };
    public static (CommandType type, int argLength) GetCommandInfo(string command)
    {
        if (commandInfo.TryGetValue(command, out var info))
        {
            return info;
        }
        return (CommandType.C_UNKNOWN, 0);
    }
}

public enum CommandType
{
    C_UNKNOWN = -1, // 初期値/エラー状態
    C_ARITHMETIC,   // 0
    C_PUSH,         // 1
    C_POP,          // 2
    C_LABEL,
    C_GOTO,
    C_IF,
    C_FUNCTION,
    C_RETURN,
    C_CALL
}
