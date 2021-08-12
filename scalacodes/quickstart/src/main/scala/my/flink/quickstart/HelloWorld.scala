package my.flink.quickstart

object HelloWorld {
  def main(args: Array[String]): Unit = {
    println("Hello, world")
    var myVar: String = "Foo"
    var myConst: String = "Foo"
    var myStr = "Foo"
    val xmax, ymax = 100
    var pa: (Int, String) = (40, "Foo")
    var pb = (40, "foo")
  }

  class Outer {

    class Inner {
      private def f() {
        println("f")
      }

      class InnerMost {
        f()
      }

    }

  }
}
