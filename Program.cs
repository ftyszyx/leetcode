using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Collections;

namespace ConsoleApp1
{
    class Program
    {
        static void Main(string[] args)
        { 
            MaxSatisfied(new int[] { 1,0,1,2,1,1,7,5 }, new int[] { 0,1,0,1,0,1,0,1}, 3);
        }

        public int HeightChecker(int[] heights)
        {
            var newlist = new List<int>();
            for (int count = 0; count < heights.Length; count++)
            {
                newlist.Add(heights[count]);
            }
            newlist.Sort((a, b) =>
            {
                return a - b;
            });
            int num = 0;
            for (int count = 0; count < heights.Length; count++)
            {
                if (heights[count] != newlist[count])
                {
                    num++;
                }
            }
            return num;
        }

        public int[] RearrangeBarcodes(int[] barcodes)
        {
            var oklist = new int[barcodes.Length];
            var removeindex = new int[barcodes.Length];
            for (int i = 0; i < barcodes.Length; i+=2)
            {
                oklist[i] = barcodes[i];
                removeindex[i] = 1;
            }
            
            for (int i = 1; i < barcodes.Length; i += 2)
            { 

                if (i < barcodes.Length - 1)
                {
                    if (oklist[i - 1] != barcodes[i] && oklist[i +1] != barcodes[i])
                    {
                        oklist[i] = barcodes[i];
                        removeindex[i] = 1;
                    }
                }
                else
                {
                    if (oklist[i - 1] != barcodes[i])
                    {
                        oklist[i] = barcodes[i];
                        removeindex[i] = 1;
                    }
                } 
            }
            return oklist;
        }

        public static int MaxSatisfied(int[] customers, int[] grumpy, int X)
        {
            var oklist = new int[customers.Length];
            var pre = new int[customers.Length];
            var preok = new int[customers.Length];
            int num = 0;
            var totalnum = 0;
            
            for (int i = 0; i < customers.Length; i++)
            {
                if (grumpy[i] == 0)
                {
                    num = num + customers[i];
                    oklist[i] = 1;
                }
                else
                {
                    oklist[i] = 0;
                }
                totalnum= totalnum + customers[i];
                if (i == 0)
                {
                    pre[i] = customers[i];
                    if (grumpy[i] == 0)
                        preok[i]= customers[i];
                }
                else
                {
                    pre[i] = pre[i - 1] + customers[i];
                    if (grumpy[i] == 0)
                        preok[i] = preok[i - 1] + customers[i];
                    else
                        preok[i] = preok[i - 1];
                }
            }
            if (customers.Length <= X)
            {
                return totalnum;
            }

            int maxaddnum = 0;
            for (int i = 0; i < customers.Length; i++)
            {
                int canaddnum = 0;
                var end = (i + X-1) >= (customers.Length - 1) ? (customers.Length - 1) : i + X-1;
                if (i == 0)
                {
                    canaddnum = pre[end]  - preok[end] ;
                    
                }
                else
                {
                    canaddnum = pre[end] - pre[i-1] - (preok[end] - preok[i - 1]);
                }
                if (canaddnum > maxaddnum)
                    maxaddnum = canaddnum;

            }


            return num+ maxaddnum;
        }
    }
}
