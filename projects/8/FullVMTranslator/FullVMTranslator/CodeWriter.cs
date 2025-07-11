using System.Text;

public interface ICodeWriter
{
    void SetFileName(string filename);
    void WriteBootstrap();
    void WriteArithmetic(string command);
    void WritePushPop(CommandType command, string segment, int index);
    void WriteLabel(string label);
    void WriteGoto(string label);
    void WriteIf(string label);
    void WriteFunction(string functionName, int nVars);
    void WriteCall(string functionName, int nArgs);
    void WriteReturn();
}

public class CodeWriter : ICodeWriter
{
    private readonly TextWriter writer;
    private string? filename;
    private int labelCounter = 0;
    private int callCounter = 0;

    public CodeWriter(TextWriter writer)
    {
        this.writer = writer;
    }

    public void SetFileName(string filename)
    {
        this.filename = filename;
    }

    public void WriteBootstrap()
    {
        writer.WriteLine(AssemblyCodeBlock.BOOTSTRAP);
        WriteCall("Sys.init", 0);
    }

    public void WriteArithmetic(string command)
    {
        string assembly = command switch
        {
            "add" => AssemblyCodeBlock.ADD,
            "sub" => AssemblyCodeBlock.SUB,
            "neg" => AssemblyCodeBlock.NEG,
            "and" => AssemblyCodeBlock.AND,
            "or" => AssemblyCodeBlock.OR,
            "not" => AssemblyCodeBlock.NOT,
            "eq" => string.Format(AssemblyCodeBlock.EQ, labelCounter++),
            "gt" => string.Format(AssemblyCodeBlock.GT, labelCounter++),
            "lt" => string.Format(AssemblyCodeBlock.LT, labelCounter++),
            _ => throw new ArgumentException($"unsupported arithmetic command: {command}")
        };
        writer.WriteLine(assembly);
    }

    public void WritePushPop(CommandType command, string segment, int index)
    {
        string assembly;
        switch (command)
        {
            case CommandType.C_PUSH:
                assembly = WritePush(segment, index);
                break;
            case CommandType.C_POP:
                assembly = WritePop(segment, index);
                break;
            default:
                throw new ArgumentException($"unsupported command type for WritePushPop: {command}");
        }
        writer.WriteLine(assembly);
    }

    public void WriteLabel(string label)
    {
        if (string.IsNullOrEmpty(filename))
        {
            throw new InvalidOperationException("filename not set for label");
        }
        string assembly = string.Format("({0}${1})", filename, label);
        writer.WriteLine(assembly);
    }

    public void WriteGoto(string label)
    {
        if (string.IsNullOrEmpty(filename))
        {
            throw new InvalidOperationException("filename not set for goto");
        }
        string assembly = string.Format(AssemblyCodeBlock.GOTO, filename, label);
        writer.WriteLine(assembly);
    }

    public void WriteIf(string label)
    {
        if (string.IsNullOrEmpty(filename))
        {
            throw new InvalidOperationException("filename not set for if-goto");
        }
        string assembly = string.Format(AssemblyCodeBlock.IF, filename, label);
        writer.WriteLine(assembly);
    }

    public void WriteFunction(string functionName, int nVars)
    {
        if (string.IsNullOrEmpty(filename))
        {
            throw new InvalidOperationException("filename not set for function");
        }

        var b = new StringBuilder();
        for (int i = 0; i < nVars; i++)
        {
            b.AppendLine(WritePush("constant", 0));
        }
        var assembly = string.Format(
            AssemblyCodeBlock.FUNCTION,
            filename,
            functionName,
            nVars,
            b.ToString().Trim());
        writer.WriteLine(assembly);
    }
    public void WriteCall(string functionName, int nArgs)
    {
        if (string.IsNullOrEmpty(filename))
        {
            throw new InvalidOperationException("filename not set for call");
        }
        var returnLabel = $"{filename}$ret.{callCounter++}";

        var assembly = string.Format(
            AssemblyCodeBlock.CALL,
            filename,
            functionName,
            nArgs,
            5 + nArgs,
            returnLabel);
        writer.WriteLine(assembly);
    }

    public void WriteReturn()
    {
        string assembly = AssemblyCodeBlock.RETURN;
        writer.WriteLine(assembly);
    }

    private string WritePush(string segment, int index)
    {
        string assembly = segment switch
        {
            "constant" => string.Format(AssemblyCodeBlock.PUSH_CONSTANT, index),
            "local" => string.Format(AssemblyCodeBlock.PUSH_LOCAL, index),
            "argument" => string.Format(AssemblyCodeBlock.PUSH_ARGUMENT, index),
            "this" => string.Format(AssemblyCodeBlock.PUSH_THIS, index),
            "that" => string.Format(AssemblyCodeBlock.PUSH_THAT, index),
            "temp" => string.Format(AssemblyCodeBlock.PUSH_TEMP, index),
            "pointer" => index is 0 or 1
                ? string.Format(AssemblyCodeBlock.PUSH_POINTER, index == 0 ? "THIS" : "THAT")
                : throw new ArgumentException($"invalid index for pointer segment: {index} (must be 0 or 1)"),
            "static" => !string.IsNullOrEmpty(filename)
                ? string.Format(AssemblyCodeBlock.PUSH_STATIC, filename, index)
                : throw new InvalidOperationException("filename not set for static variable"),
            _ => throw new ArgumentException($"unsupported segment for push command: {segment}")
        };
        return assembly;
    }

    private string WritePop(string segment, int index)
    {
        string assembly = segment switch
        {
            "local" => string.Format(AssemblyCodeBlock.POP_LOCAL, index),
            "argument" => string.Format(AssemblyCodeBlock.POP_ARGUMENT, index),
            "this" => string.Format(AssemblyCodeBlock.POP_THIS, index),
            "that" => string.Format(AssemblyCodeBlock.POP_THAT, index),
            "temp" => string.Format(AssemblyCodeBlock.POP_TEMP, index),
            "pointer" => index is 0 or 1
                ? string.Format(AssemblyCodeBlock.POP_POINTER, index == 0 ? "THIS" : "THAT")
                : throw new ArgumentException($"invalid index for pointer segment: {index} (must be 0 or 1)"),
            "static" => !string.IsNullOrEmpty(filename)
                ? string.Format(AssemblyCodeBlock.POP_STATIC, filename, index)
                : throw new InvalidOperationException("filename not set for static variable"),
            _ => throw new ArgumentException($"unsupported segment for pop command: {segment}")
        };
        return assembly;
    }
}

class AssemblyCodeBlock
{
    public const string BOOTSTRAP = @"// Bootstrap code
@256
D=A
@SP
M=D
";
    public const string ADD = @"// add
@SP
M=M-1 // SP--
A=M   // D = *SP
D=M
@SP
M=M-1 // SP--
A=M
M=D+M
@SP
M=M+1 // SP++
";
    public const string SUB = @"// sub
@SP
M=M-1 // SP--
A=M   // D = *SP
D=M
@SP
M=M-1 // SP--
A=M
M=M-D
@SP
M=M+1 // SP++
";
    public const string NEG = @"// neg
@SP
M=M-1 // SP--
A=M
M=-M
@SP
M=M+1 // SP++
";
    public const string AND = @"// and
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=M&D
@SP
M=M+1
";
    public const string OR = @"// or
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=D|M
@SP
M=M+1
";
    public const string NOT = @"// not
@SP
M=M-1
A=M
M=!M
@SP
M=M+1
";

    public const string EQ = @"// eq
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@EQ_TRUE_{0}
D;JEQ
@SP
A=M
M=0
@EQ_END_{0}
0;JMP
(EQ_TRUE_{0})
@SP
A=M
M=-1
(EQ_END_{0})
@SP
M=M+1
";
    public const string GT = @"// gt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@GT_TRUE_{0}
D;JGT
@SP
A=M
M=0
@GT_END_{0}
0;JMP
(GT_TRUE_{0})
@SP
A=M
M=-1
(GT_END_{0})
@SP
M=M+1
";

    public const string LT = @"// lt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
D=M-D
@LT_TRUE_{0}
D;JLT
@SP
A=M
M=0
@LT_END_{0}
0;JMP
(LT_TRUE_{0})
@SP
A=M
M=-1
(LT_END_{0})
@SP
M=M+1
";

    public const string LABEL = @"// label {1} from {0}
({0}${1})
";
    public const string GOTO = @"// goto {1} from {0}
@{0}${1}
0;JMP
";
    public const string IF = @"// if-goto {1} from {0}
@SP
M=M-1 // SP--
A=M   // A = *SP
D=M   // D = *SP
@{0}${1} // {0}${1} is a label
D;JNE // if D != 0, jump to label {0}${1}
";

    public const string FUNCTION = @"// function {1}.  nVar={2} from {0}
({1})

// nVar={2}だけ0で初期化
{3}
";

    /// <summary>
    /// {0} is the current filename
    /// {1} is the function name
    /// {2} is the number of arguments,
    /// {3} is the number of arguments to be passed to the function: 5 + nArgs
    /// {4} is the return address label
    /// </summary>
    public const string CALL = @"// call {1} with nArgs={2} from {0}
// push return address {4}
@{4}
D=A
@SP
A=M
M=D
@SP
M=M+1 // SP++

// push LCL
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

// push ARG
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

// push THIS
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

// push THAT
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1

// ARG = SP - nArgs - 5
@SP
D=M
@{3}
D=D-A
@ARG
M=D

// LCL = SP
@SP
D=M
@LCL
M=D

// goto {1} from {0}
@{1}
0;JMP

// return address label
({4})
";

    public const string RETURN = @"// return
// frame = LCL
@LCL
D=M
@R13
M=D

// return adress = *(frame - 5)
@5
A=D-A
D=M
@R14
M=D   // R14 = return adress

// *ARG = pop()
@SP
M=M-1 // SP--
A=M
D=M
@ARG
A=M
M=D

// SP = ARG + 1
@ARG
D=M
@SP
M=D+1

// THAT = *(frame - 1)
@R13
M=M-1
A=M
D=M
@THAT
M=D

// THIS = *(frame - 2)
@R13
M=M-1
A=M
D=M
@THIS
M=D

// ARG = *(frame - 3)
@R13
M=M-1
A=M
D=M
@ARG
M=D

// LCL = *(frame - 4)
@R13
M=M-1
A=M
D=M
@LCL
M=D

// goto return address
@R14
A=M
0;JMP 
";

    public const string PUSH_CONSTANT = @"// push constant {0}
@{0}
D=A
@SP
A=M
M=D
@SP
M=M+1
";
    public const string PUSH_LOCAL = @"// push local {0}
@LCL
D=M
@{0}
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
";
    public const string PUSH_ARGUMENT = @"// push argument {0}
@ARG
D=M
@{0}
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
";
    public const string PUSH_THIS = @"// push this {0}
@THIS
D=M
@{0}
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
";
    public const string PUSH_THAT = @"// push that {0}
@THAT
D=M
@{0}
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
";
    public const string PUSH_TEMP = @"// push temp {0}
@5
D=A
@{0}
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
";
    public const string PUSH_POINTER = @"// push pointer {0}
@{0}
D=M
@SP
A=M
M=D
@SP
M=M+1
";
    public const string PUSH_STATIC = @"// push static {1}
@{0}.{1} // アセンブラによってアドレス16から自動で割り当てられる
D=M
@SP
A=M
M=D
@SP
M=M+1
";

    public const string POP_LOCAL = @"// pop local {0}
@LCL
D=M
@{0}
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
";
    public const string POP_ARGUMENT = @"// pop argument {0}
@ARG
D=M
@{0}
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
";
    public const string POP_THIS = @"// pop this {0}
@THIS
D=M
@{0}
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
";
    public const string POP_THAT = @"// pop that {0}
@THAT
D=M
@{0}
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
";
    public const string POP_TEMP = @"// pop temp {0}
@5
D=A
@{0}
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
";
    public const string POP_POINTER = @"// pop pointer {0}
@SP
M=M-1
A=M
D=M
@{0}
M=D
";
    public const string POP_STATIC = @"// pop static {1}
@SP
M=M-1
A=M
D=M
@{0}.{1}
M=D
";
}
