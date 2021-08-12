package com.demo.flink;

import com.alibaba.fastjson.JSON;
import com.demo.flink.model.Student;
import com.demo.flink.sink.SinkToMySQL;
import org.apache.flink.api.common.functions.FilterFunction;
import org.apache.flink.api.common.functions.FlatMapFunction;
import org.apache.flink.api.common.functions.MapFunction;
import org.apache.flink.api.common.functions.ReduceFunction;
import org.apache.flink.api.common.serialization.SimpleStringSchema;
import org.apache.flink.api.java.functions.KeySelector;
import org.apache.flink.streaming.api.datastream.KeyedStream;
import org.apache.flink.streaming.api.datastream.SingleOutputStreamOperator;
import org.apache.flink.streaming.api.environment.StreamExecutionEnvironment;
import org.apache.flink.streaming.connectors.kafka.FlinkKafkaConsumer;
import org.apache.flink.util.Collector;

import java.util.Properties;

public class FlinkTransform {
    public static void main(String[] args) throws Exception {
        final StreamExecutionEnvironment env = StreamExecutionEnvironment.getExecutionEnvironment();

        Properties props = new Properties();
        props.put("bootstrap.servers", "192.168.2.122:32400");
        props.put("zookeeper.connect", "192.168.2.122:2181");
        props.put("group.id", "metric-group");
        props.put("key.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
        props.put("value.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
        props.put("auto.offset.reset", "latest");

        SingleOutputStreamOperator<Student> student = env.addSource(new FlinkKafkaConsumer<>(
                        "student",   //这个 kafka topic 需要和上面的工具类的 topic 一致
                        new SimpleStringSchema(),
                        props)).setParallelism(1)
                .map(string -> JSON.parseObject(string, Student.class)); //Fastjson 解析字符串成 student 对象

//        SingleOutputStreamOperator<Student> map = student.map(new MapFunction<Student, Student>() {
//            @Override
//            public Student map(Student value) throws Exception {
//                Student s1 = new Student();
//                s1.setId(value.getId());
//                s1.setName(value.getName());
//                s1.setPassword(value.getPassword());
//                s1.setAge(value.getAge() + 5);
//                return s1;
//            }
//        });
//        map.print();

//        SingleOutputStreamOperator<Student> flatMap = student.flatMap(new FlatMapFunction<Student, Student>() {
//            @Override
//            public void flatMap(Student value, Collector<Student> out) throws Exception {
//                if (value.getId() % 2 == 0) {
//                    out.collect(value);
//                }
//            }
//        });
//        flatMap.print();
//
//        SingleOutputStreamOperator<Student> filter = student.filter(new FilterFunction<Student>() {
//            @Override
//            public boolean filter(Student value) throws Exception {
//                if (value.getId() > 95) {
//                    return true;
//                }
//                return false;
//            }
//        });
//        filter.print();

//        KeyedStream<Student, Integer> keyBy = student.keyBy(new KeySelector<Student, Integer>() {
//            @Override
//            public Integer getKey(Student value) throws Exception {
//                return value.getAge();
//            }
//        });
//        keyBy.print();

//        SingleOutputStreamOperator<Student> reduce = student.keyBy(new KeySelector<Student, Integer>() {
//            @Override
//            public Integer getKey(Student value) throws Exception {
//                return value.getAge();
//            }
//        }).reduce(new ReduceFunction<Student>() {
//            @Override
//            public Student reduce(Student value1, Student value2) throws Exception {
//                Student student1 = new Student();
//                student1.name = value1.name + value2.name;
//                student1.id = (value1.id + value2.id) / 2;
//                student1.password = value1.password + value2.password;
//                student1.age = (value1.age + value2.age) / 2;
//                return student1;
//            }
//        });
//        reduce.print();


//        KeyedStream<Student, Integer> keyBy = student.keyBy(new KeySelector<Student, Integer>() {
//            @Override
//            public Integer getKey(Student value) throws Exception {
//                return value.getAge();
//            }
//        });
//        keyBy.sum(0);
//        keyBy.sum("key");
//        keyBy.min(0);
//        keyBy.min("key");
//        keyBy.max(0);
//        keyBy.max("key");
//        keyBy.minBy(0);
//        keyBy.minBy("key");
//        keyBy.maxBy(0);
//        keyBy.maxBy("key");

        env.execute("Flink add transform");
    }
}