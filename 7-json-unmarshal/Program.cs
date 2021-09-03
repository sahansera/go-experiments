using System;
using System.Text.Json;

public class T
{
    public string HelloWorld { get; set; }
}

public class Program
{
	public static void Main()
	{
		string jsonString = @"{""HelloWorld"": ""1"", ""helloWorld"": ""2""}";
		Console.WriteLine(JsonSerializer.Deserialize<T>(jsonString).HelloWorld);
	}
}