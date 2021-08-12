package my.flink.quickstart;

import org.apache.flink.api.common.functions.MapFunction;

public class MyMapFunction<T, O> implements MapFunction<T, O> {
    @Override
    public O map(T t) throws Exception {
        return null;
    }
}
