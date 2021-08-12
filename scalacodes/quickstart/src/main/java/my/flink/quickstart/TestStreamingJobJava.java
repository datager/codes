package my.flink.quickstart;

import org.apache.flink.streaming.api.datastream.DataStream;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;

import java.util.Arrays;

public class TestStreamingJobJava {
    public static void main(String[] args) throws Exception {
        StreamExecutionEnvironment env = StreamExecutionEnvironment.getExecutionEnvironment();
//        env.getConfig().enable
//        DataStream<String> input = env.fromElements("aaa", "bbb");
//        DataStream<Integer> s = input.map(new MyMapFunction<String, Integer>())
//                .returns(new TypeHint<Integer>() {
//        });

        String[] e = new String[]{"hello", "flink"};
        DataStream<String> s = env.fromCollection(Arrays.asList(e));
        s.print(">>>");
        env.execute("javaJob");
    }
}
