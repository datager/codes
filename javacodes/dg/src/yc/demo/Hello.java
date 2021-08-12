package yc.demo;

import com.sun.tools.javac.util.List;

import java.util.Collection;
import java.util.HashMap;
import java.util.HashSet;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class Hello {
    public static void main(String[] args) {
        System.out.println("Hello world!");


//        Integer i1 = 127;
//        Integer i2 = 127;
//        System.err.println(i1 == i2);
//        System.err.println(i1.equals(i2));
//
//        i1 = 128;
//        i2 = 128;
//        System.err.println(i1 == i2);
//        System.err.println(i1.equals(i2));


        Integer i1 = 6;
        Integer i2 = 6;
        System.out.println((i1 == i2));//true


        Integer i3 = new Integer(6);
        Integer i4 = new Integer(6);

        System.out.println((6 == i3));
        System.out.println((128 == i3));
        System.out.println((i4 == 128));

        System.out.println((i3 == i4) + " " + i3.hashCode() + " " + i4.hashCode());

        List<String> simpleList = List.of("hello", "world");

        String paylod = IntStream.rangeClosed(1, 100).mapToObj(__ -> "a").collect(Collectors.joining(""));
        System.out.println(paylod);
    }
}
