package yc.demo;


public class Singleton {
    private Singleton() {
    }

    public static Singleton getSingleton() {
        return Holder.singleton;
    }

    private static class Holder {
        private static Singleton singleton = new Singleton();
    }
}
