using System;
using System.IO;

namespace BasicAssembler;

class Program
{
    static void Main(string[] args)
    {
        if (args.Length != 1)
        {
            Console.WriteLine("Usage: BasicAssembler <inputfile.asm>");
            return;
        }

        var inputFilePath = args[0];

        if (!File.Exists(inputFilePath))
        {
            Console.WriteLine($"Error: File '{inputFilePath}' not found.");
            return;
        }

        var outputFilePath = Path.ChangeExtension(inputFilePath, ".hack");
        using var outputFile = new StreamWriter(outputFilePath);

        IParser parser = Parser.CreateParser(inputFilePath);
        ICode code = new Code();

        var assembler = new Assembler(code, parser);
        assembler.Assemble(outputFile);

        Console.WriteLine($"Successfully assembled to '{outputFilePath}'.");
    }
}