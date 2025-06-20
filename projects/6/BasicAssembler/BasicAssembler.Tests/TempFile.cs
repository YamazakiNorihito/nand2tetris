namespace BasicAssembler.Tests;

using System;
using System.IO;

/// <summary>
/// テスト用一時ファイルヘルパー。Dispose時に自動削除されます。
/// </summary>
public sealed class TempFile : IDisposable
{
    public string Path { get; }

    public TempFile(string[] lines, string? fileName = null)
    {
        if (string.IsNullOrEmpty(fileName))
        {
            Path = System.IO.Path.GetTempFileName();
        }
        else
        {
            Path = fileName;
        }
        File.WriteAllLines(Path, lines);
    }

    public void Dispose()
    {
        if (File.Exists(Path))
            File.Delete(Path);
    }
}