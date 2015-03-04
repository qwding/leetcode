import java.util.Collections;
import java.util.Iterator;
import java.util.List;
import java.util.LinkedList;
import java.util.Collections;
import java.util.Comparator;

public class Solution {
	public static List<Interval> merge(List<Interval> intervals){
		Collections.sort(intervals,new Comparator<Interval>(){
			public int compare(Interval intervalOne,Interval intervalTwo){
				return intervalOne.start - intervalTwo.start;
			}
		});
		int size = intervals.size();
		for(int i = 0; i < size; ++i){
			//System.out.println("this is : start " + intervals.get(i).start + " end "+ intervals.get(i).end);
			while(i < size && (i+1) < size && intervals.get(i).end >= intervals.get(i+1).start){
				//System.out.println("this is : start " + intervals.get(i+1).start + " end "+ intervals.get(i+1).end);
				intervals.get(i).end = intervals.get(i).end <intervals.get(i+1).end?intervals.get(i+1).end:intervals.get(i).end;
				intervals.remove(i+1);
				--size;
			}
		}
		return intervals;
	}
	public static void main(String[] args){
		Interval in1 = new Interval(1,3);
		Interval in2 = new Interval(3,6);
		Interval in3 = new Interval(4,22);
		Interval in4 = new Interval(8,10);
		Interval in5 = new Interval(15,18);
		
		List list = new LinkedList();
		list.add(in1);
		list.add(in2);
		list.add(in3);
		list.add(in4);
		list.add(in5);
		merge(list);
		System.out.println("---------------------");
		Iterator<Interval> iter = list.iterator();
		while(iter.hasNext()){
			Interval tempOne = iter.next(); 
			System.out.println(" this is " + tempOne.start + " " + tempOne.end);
		}
	}
}


public class Interval {
    int start;
    int end;
    Interval() { start = 0; end = 0; }
    Interval(int s, int e) { start = s; end = e; }
} 