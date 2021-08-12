package my.flink.quickstart

import org.apache.flink.api.java.utils.ParameterTool
import org.apache.flink.api.scala._
import org.apache.flink.streaming.api.scala.StreamExecutionEnvironment


object WordCount {

  def main(args: Array[String]) {
    if (args.length != 1) {
      println("Please give input parameter.")
      System.exit(1)
    }
    //    val env = ExecutionEnvironment.getExecutionEnvironment
    //    val text = env.readTextFile(args(0))
    //    val counts = text.flatMap { _.toLowerCase.split("\\W+") filter { _.nonEmpty } }
    //      .map { (_, 1) }
    //      .groupBy(0)
    //      .sum(1)
    //    counts.print()

    //    val env = ExecutionEnvironment.getExecutionEnvironment
    //    val dataSet = env.fromElements(("hello", 1), ("flink, 3"))
    //    val g: GroupedDataSet[(String, Int)] = dataSet.groupBy(0)
    //    g.max(1)

    val env = StreamExecutionEnvironment.getExecutionEnvironment
    //    val p = env.fromElements(("Alex", 18), ("Peter", 43))
    //    p.keyBy("_1").sum("_2").print()

    env.execute("")
  }

  class CompelexClass(var nested: NestedClass, var tag: String) {
    def this() {
      this(null, "")
    }
  }

  class NestedClass(
                     var id: Int,
                     tuple: (Long, Long, String)) {
    def this() {
      this(0, (0, 0, ""))
    }
  }

}
