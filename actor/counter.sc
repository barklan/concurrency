import scala.actors.Actor
import scala.actors.Actor._

case class Inc(amount: Int)
case class Value

class Counter extends Actor {
    var counter: Int = 0;

    def act() = {
        while (true) {
            receive {
                case Inc(amount) =>
                    counter += amount
                case Value =>
                    println("Value is "+counter)
                    exit()
            }
        }
    }
}

object ActorTest extends Application {
    val counter = new Counter
    counter.start()

    for (i <- 0 until 100000) {
        counter ! Inc(1)
    }
    counter ! Value
    // Output: Value is 100000
}
