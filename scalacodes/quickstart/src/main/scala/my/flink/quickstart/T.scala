package my.flink.quickstart

import org.apache.flink.streaming.api.scala._

object T {

  case class WordCount(word: String, count: Int)

  def main(args: Array[String]): Unit = {

    val env = StreamExecutionEnvironment.getExecutionEnvironment
    val s = env.fromElements(Tuple2(1, 3), Tuple2("a", "b"), Tuple2(5, "c"))
    s.print("b")

    // POJO
    //    val env = StreamExecutionEnvironment.getExecutionEnvironment
    //    val s = env.fromElements(new PersonJava("Peter", 14), new PersonJava("Linda", 25))
    //    s.keyBy("name").print("name")


    //    val input1 = env.fromElements("1", "2")
    //    input1.print("111")
    //    val input = env.fromElements(WordCount("hello", 1), WordCount("world", 2))
    //    //    val keyStream1 = input.keyBy("word")
    //    //    val keyStream2 = input.keyBy(0)
    //    //    keyStream1.print("1")
    //    //    keyStream2.print("2")
    //    input.print("1")
    env.execute("job")
  }
}

class PersonScala(var name: String, var age: Int) {
  def this() {
    this(null, -1)
  }
}


