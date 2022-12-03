class Program
{
    static void Main(string[] args)
    {
        // 'scores' is a list of scores for each elf.
        List<int> scores = new List<int>();
        // 'counter' contains the score of the current elf.
        int counter = 0;
        // 'value' is a temporary variable used to store the result of parsing the number in each line.
        int value = 0;

        // for each line inside the input file "input":
        foreach (string line in System.IO.File.ReadLines(@"input"))
        {
            // check for empty lines first; when we get to an empty line we need to store the score
            // of this elf and then continue to the next line.
            if (line == "")
            {
                scores.Add(counter);
                counter = 0;
                continue;
            }

            // If we get here it means that the line wasn't empty, so we first try to parse the text
            // into an 'int'; this is the number of calories from the current line.
            try
            {
                value = Int32.Parse(line);
            }
            catch (FormatException)
            {
                // if we get here it means that we couldn't parse the number in the current line.
                Console.WriteLine($"cannot parse '{line}'");
                return;
            }

            // add the calories to the counter.
            counter += value;
        }

        // we sort the list of calories of each elf, so that our list starts with the lowest number
        // of calories and ends with the highest.
        scores.Sort();
        Console.WriteLine($"first score: {scores[0]}, last score: {scores[scores.Count-1]}");

        // Now we print the number of calories of the three elves with most calories.
        int top3 = 0;
        for (int i = scores.Count-3; i < scores.Count; i++)
        {
            Console.WriteLine(scores[i]);
            top3 += scores[i];
        }

        Console.WriteLine($"the sum of the calories of the top three elves is: {top3}");
    }
}
