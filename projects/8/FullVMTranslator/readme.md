


```bash
dotnet --version

dotnet new sln -n BasicAssembler

dotnet new console -n BasicAssembler
dotnet sln add BasicAssembler/BasicAssembler.csproj

dotnet new xunit -n BasicAssembler.Tests
dotnet sln add BasicAssembler.Tests/BasicAssembler.Tests.csproj

dotnet build

dotnet run --project BasicAssembler
dotnet test


```