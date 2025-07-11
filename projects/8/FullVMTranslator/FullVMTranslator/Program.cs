class Program
{
    private static FileStreamOptions _readFileStreamOptions = new FileStreamOptions
    {
        Access = FileAccess.Read,
        Mode = FileMode.Open,
        Share = FileShare.Read
    };

    private static FileStreamOptions _writeFileStreamOptions = new FileStreamOptions
    {
        Access = FileAccess.Write,
        Mode = FileMode.Create,
        Share = FileShare.None
    };

    static void Main(string[] args)
    {
        if (args.Length != 1)
        {
            Console.Error.WriteLine("Usage: VMTranslator <source.vm|source_directory>");
            Environment.Exit(1);
        }

        string inputFilePath = args[0];
        var (vmFiles, outputFilePath) = GetVmFilesAndOutputFilePath(inputFilePath);

        try
        {
            Translation(vmFiles, outputFilePath);
            Console.WriteLine($"Translation complete.");
        }
        catch (Exception ex)
        {
            Console.Error.WriteLine($"Error: {ex.Message}");
            Environment.Exit(1);
        }
    }

    private static void Translation(List<string> vmFiles, string outputFilePath)
    {
        using var writer = new StreamWriter(outputFilePath, _writeFileStreamOptions);

        var codeWriter = new CodeWriter(writer);
        if (vmFiles.Any(f => f.Contains("Sys.vm", StringComparison.CurrentCultureIgnoreCase)))
        {
            codeWriter.SetFileName("Bootstrap");
            codeWriter.WriteBootstrap();
        }

        foreach (var vmFile in vmFiles)
        {
            string fileName = Path.GetFileNameWithoutExtension(vmFile);
            codeWriter.SetFileName(fileName);
            if (vmFile.Contains("Sys.vm", StringComparison.CurrentCultureIgnoreCase))
            {
                codeWriter.WriteBootstrap();
            }

            using var reader = new StreamReader(vmFile, _readFileStreamOptions);
            var parser = new Parser(reader);
            while (parser.HasMoreLines())
            {
                parser.Advance();
                switch (parser.CommandType)
                {
                    case CommandType.C_ARITHMETIC:
                        codeWriter.WriteArithmetic(parser.Arg1);
                        break;
                    case CommandType.C_PUSH:
                    case CommandType.C_POP:
                        codeWriter.WritePushPop(parser.CommandType, parser.Arg1, parser.Arg2);
                        break;
                    case CommandType.C_LABEL:
                        codeWriter.WriteLabel(parser.Arg1);
                        break;
                    case CommandType.C_GOTO:
                        codeWriter.WriteGoto(parser.Arg1);
                        break;
                    case CommandType.C_IF:
                        codeWriter.WriteIf(parser.Arg1);
                        break;
                    case CommandType.C_FUNCTION:
                        codeWriter.WriteFunction(parser.Arg1, parser.Arg2);
                        break;
                    case CommandType.C_CALL:
                        codeWriter.WriteCall(parser.Arg1, parser.Arg2);
                        break;
                    case CommandType.C_RETURN:
                        codeWriter.WriteReturn();
                        break;
                    default:
                        Console.Error.WriteLine($"Unknown command type: {parser.CommandType}");
                        Environment.Exit(1);
                        break;
                }
            }
        }
    }

    private static (List<string> vmFiles, string outputFilePath) GetVmFilesAndOutputFilePath(string inputPath)
    {
        if (Directory.Exists(inputPath))
        {
            string dirName = new DirectoryInfo(inputPath).Name;
            var outputFilePath = Path.Combine(inputPath, dirName + ".asm");

            var vmFiles = Directory.GetFiles(inputPath, "*.vm").ToList();
            return (vmFiles, outputFilePath);
        }
        else if (File.Exists(inputPath))
        {
            var outputFilePath = Path.ChangeExtension(inputPath, ".asm");
            var vmFiles = new List<string> { inputPath };
            return (vmFiles, outputFilePath);
        }
        else
        {
            throw new FileNotFoundException($"Input file or directory not found: {inputPath}");
        }
    }
}