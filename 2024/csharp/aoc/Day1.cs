using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace aoc;
public class Day1
{
    public void Challenge1()
    {
        var input_file = "day1_input-1.csv";

        var lcol = new List<int>();
        var rcol = new List<int>();

        using (var filestream = File.OpenRead(input_file))
        {
            using (var stream = new StreamReader(filestream, Encoding.UTF8, true))
            {
                String? line;

                while ((line = stream.ReadLine()) != null)
                {

                }
            }

        }



    }
}
