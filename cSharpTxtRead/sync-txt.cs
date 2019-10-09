using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;

namespace cSharpTxtRead
{
    class Program
    {
        static Dictionary<string, int> totalWords = new Dictionary<string, int>();
        static void Main(string[] args)
        {
            System.Diagnostics.Stopwatch stopwatch = new System.Diagnostics.Stopwatch();
            stopwatch.Start();
            object locker = new object();
            var dir = "C:\\Users\\mehmetakif.tatar\\Desktop\\Proje\\txtReadPerform\\txt-repo";
            var x = Directory.GetFiles(dir);



            foreach (var v in x) // Uncomment for the syncrinos execution
            {
                string text = System.IO.File.ReadAllText(v);
                var split = text.Split(" ");

                foreach (var item in split)
                {

                    if (totalWords.ContainsKey(item))
                    {
                        totalWords[item] = totalWords[item] + 1;
                    }
                    else
                    {
                        totalWords[item] = 1;
                    }
                }
            }
            stopwatch.Stop();
            System.Console.WriteLine("sync");
            System.Console.WriteLine(stopwatch.ElapsedMilliseconds / 1000 + "s");
            System.Console.WriteLine(totalWords["is"]);

        }
    }
}
