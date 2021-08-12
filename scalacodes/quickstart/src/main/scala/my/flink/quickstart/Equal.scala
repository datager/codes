package my.flink.quickstart

import scala.util.matching.Regex

trait Equal {
  def isEqual(x: Any): Boolean

  def isNotEqual(x: Any): Boolean = !isEqual(x)
}

class Point(xc: Int, yc: Int) extends Equal {
  var x: Int = xc
  var y: Int = yc

  override def isEqual(obj: Any): Boolean = {
    obj.isInstanceOf[Point] && obj.asInstanceOf[Point].x == x
  }

  override def isNotEqual(obj: Any): Boolean = {
    !isEqual(obj)
  }
}

object Test1 {
  def main(args: Array[String]) {

    val x = Test1(5)
    println(x)

    x match
    {
      case Test1(num) => println(x + " 是 " + num + " 的两倍！")
      //unapply 被调用
      case _ => println("无法计算")
    }

  }
  def apply(x: Int) = x*2
  def unapply(z: Int): Option[Int] = if (z%2==0) Some(z/2) else None
}

//
//  case class Person(name: String, age: Int)
//
//  def main(args: Array[String]): Unit = {
//
//    val p = new Regex("(S|s)cala")
//    val str = "Scala is scalable and cool"
//    println(p findFirstIn str)
//
//
//    val a = new Person("a", 1)
//    val b = Person("b", 2)
//    val c = Person("c", 3)
//
//    for (p <- List(a, b, c)) {
//      p match {
//        case Person("a", 1) => println("a")
//        case Person("b", 2) => println("b")
//        case Person(name, ag) => println("name" + name + "age" + ag)
//      }
//    }
//
//    //    val p1 = new Point(2, 3)
//    //    val p2 = new Point(2, 4)
//    //    val p3 = new Point(3, 3)
//    //
//    //    println(p1.isEqual(p2))
//    //    //    println(p1.isEqual(p3))
//    //    //    println(p1.isEqual(2))
//    //
//    //    println(matchTest(3))
//    //  }
//    //
//    //  def matchTest(x: Int): String = x match {
//    //    case 1 => "1"
//    //    case 2 => "two"
//    //    case _ => "many"
//  }
