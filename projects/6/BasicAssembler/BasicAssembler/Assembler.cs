namespace BasicAssembler;

public class Assembler
{
    private readonly ICode _code;
    private readonly IParser _parser;

    public Assembler(ICode code, IParser parser)
    {
        _code = code;
        _parser = parser;
    }

    public void Assemble(StreamWriter writer)
    {
        while (_parser.HasMoreLines())
        {
            _parser.Advance();
            var instructionType = _parser.InstructionType();

            if (instructionType == IParser.A_INSTRUCTION)
            {
                var symbol = _parser.Symbol();
                var value = Convert.ToInt16(symbol);
                var binary = Convert.ToString(value, 2).PadLeft(16, '0');
                writer.WriteLine(binary);
            }
            else if (instructionType == IParser.C_INSTRUCTION)
            {
                var dest = _parser.Dest();
                var comp = _parser.Comp();
                var jump = _parser.Jump();
                var binary = "111" + _code.Comp(comp) + _code.Dest(dest) + _code.Jump(jump);
                writer.WriteLine(binary);
            }
            else if (instructionType == IParser.L_INSTRUCTION)
            {
                var symbol = _parser.Symbol();
            }
        }
    }
}