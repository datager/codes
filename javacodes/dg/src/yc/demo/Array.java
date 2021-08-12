package yc.demo;

import java.util.ArrayList;

public class Array {
    public static void main(String[] args) {
        int[] a = new int[3];
        a[2] = 10;

        ArrayList<Integer> users = new ArrayList<>(1);
        for (int i = 0; i < 10000; i++) {
            users.add(1);
            System.out.println(i);
        }

        ArrayList<ArrayList<Integer>> arrays = new ArrayList<>();
        arrays.add(new ArrayList<>(1));
        System.out.println(arrays);
    }
}
