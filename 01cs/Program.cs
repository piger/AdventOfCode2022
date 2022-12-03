namespace AdventOfCode
{
    class Program
    {
        static void Main(string[] args)
        {
            List<int> scores = new List<int>();
            int counter = 0;
            int value = 0;

            foreach (string line in System.IO.File.ReadLines(@"input"))
            {
                // check for empty lines first.
                if (line == "")
                {
                    scores.Add(counter);
                    counter = 0;
                    continue;
                }

                // try to parse the number in the current line.
                try
                {
                    value = Int32.Parse(line);
                }
                catch (FormatException)
                {
                    Console.WriteLine($"cannot parse '{line}'");
                    return;
                }

                // add to the score.
                counter += value;
            }

            scores.Sort();
            Console.WriteLine($"first score: {scores[0]}, last score: {scores[scores.Count-1]}");

            for (int i = scores.Count-3; i < scores.Count; i++)
            {
                Console.WriteLine(scores[i]);
            }
        }
    }
}
