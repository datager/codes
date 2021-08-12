package yc.demo;

import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.atomic.AtomicInteger;

public class CasLock {
    private static final CountDownLatch latch = new CountDownLatch(5);
    private static AtomicInteger i = new AtomicInteger(0);
    private static int p = 0;

    public static void main(String[] args) throws InterruptedException {
        long time = System.currentTimeMillis();
        ExecutorService pool = Executors.newFixedThreadPool(5);
        for (int j = 0; j < 5; j++) {
            pool.execute(new Runnable() {
                public void run() {
                    for (int k = 0; k < 10000; k++) {
                        p++;                //不是原子操作
                        i.getAndIncrement();//调用原子类加1
                    }
                    latch.countDown();
                }
            });
        }
        latch.await();//保证所有子线程执行完成
        System.out.println(System.currentTimeMillis() - time);
        System.out.println("p=" + p);
        System.out.println("i=" + i);
        pool.shutdown();
    }
}