# Concurrency Patterns

## Introduction (up to pigeons)

// first slide with crosswalk

At first some background. So if you look around in the world at large, what you see is a lot of independently executing things. You see people in the audience doing their own thing, there's people outside, there's cars going by. All of those things are independent agents, if you will.

// second slide with people in some arena

И все эти агенты общаются между собой и действуют в единой среде.

// slide with elevators and with pigeons

Причем состояние в котором окажутся эти агенты может быть неопределена или зависеть от множества факторов вне нашего контроля.

// slide with one-way road sign

И если вы думаете о написании компьютерной программы, если вы хотите моделировать эту среду или взаимодействовать с ней - единое последовательное выполнение - не очень хороший подход.

Поэтому люди стали придумывать способы моделировать такие системы, но прежде нужно сделать важное замечание,

// slide concurrency is not parallelism

### Concurrency is not parallelism

And so cuncurrency is really a way of writing or structuring you program to deal with real world. And what I mean by concurrency is the composition of independently executing computations (usually functions, but they don't have to be). Concurrency is a way to structure software. It is not parallelism.

// slide with cpu and threads

Выполнение может быть не параллельным, но при этом конкурентным, как например на системе с одним процессором - в один момент времени процессор выплняет одну инструкцию, но при этом эти инструкции чередуются.

// slide with airport 1 and 2

### Airport example

Хорошим примером может служить аэропорт с одной взлетно посадочной полосой. Где наша сисема конкурентна, т.к. представлена множеством самолётов. Но при этом не параллельна - только один самолет может использовать полосу. Но такая канкурентная система при этом очень эффективна - в большинстве аэропортов используется только одна полоса. Но при этом остальные самолеты не стоят - они готовятся к тому чтобы использовать эту полосу.

// slide with parallel plane landing

С другой стороны, хорошо написанная параллельная программа может эффективно работать параллельно в многопроцессорной или распределенной среде.

Это свойство важно и источник заблуждений, потому что исторически люди сначала задумывались как бы сделать свою линейную программу быстрее с помощью декомпозиции на части (т.е. сделать программу конкурентной) и запустить на нескольких процессорах. Поэтому часто конкурентность (метод декомпозиции и распределения) отожествлялся с параллелизмом (т.е. непосредственным выполнением нескольких вещей в один момент времени). Только пожже люди начали задумываться над более широким применением конкурентности как методом моделирования программ. Но сначало, зачем вообще нужно распараллеливание вычислений - почему мы не можем просто увеличивать мощность.

// slide with moore'd law

### But why do we even need parallel?

В последнее время, чтобы получить возможность задействовать на практике ту дополнительную вычислительную мощность, которую предсказывает закон Мура, стало необходимо задействовать параллельные вычисления. На протяжении многих лет производители процессоров постоянно увеличивали тактовую частоту и параллелизм на уровне инструкций, так что на новых процессорах старые однопоточные приложения исполнялись быстрее без каких-либо изменений в программном коде.[8] Примерно с середины десятилетия 2000-х годов по разным причинам производители процессоров предпочитают многоядерные архитектуры, и для получения всей выгоды от возросшей производительности ЦП программы должны переписываться в соответствующей манере. Однако не каждый алгоритм поддается распараллеливанию, определяя, таким образом, фундаментальный предел эффективности решения вычислительной задачи согласно закону Амдала.

Шкала у этого графика логарифмическая, а так это вообще экспонента

// slide with exponential

// slide with imb supercomputers

Ок, первая причина конкурентного моделирования - это чтобы потом распараллелить нашу программу. Хочу уже написать что-нибудь. Но,

// slide with all languages

### Languages supporting concurrent programming

Все эти языки в той-или иной форме имеют способы конкурентного разделения задач.
Но конкурентность можно смоделировать на разных уровнях и благо для нашей цели подойдет системный уровень - во многих системах уже реализованы методы разделения.

// slide with system-level classification of concurrency

### Concurrency at system level

At the operating system level:

- Threads
- Processes

// slide with threading is better than multiprocessing

Для программ с узким назначением многопоточность лучше чем множество процессов.

// slide with items per process and per thread

Обяснение этому простое - процесс это более высокая абстракция на потоком.

**Процесс**

На уровне ядра процесс содержит один или несколько потоков, которые совместно используют ресурсы процесса, такие как память и дескрипторы файлов: процесс - это единица ресурсов, а поток - это единица планирования и выполнения. Планирование ядра обычно выполняется заранее или, что реже, совместно. Процессы обычно являются многозадачными с вытеснением, а переключение процессов является относительно дорогостоящим, помимо базовой стоимости переключения контекста.

Пример разделения процессов на about:processes в firefox.

**Поток (thread)**

Поток - это «облегченная» единица планирования ядра. В каждом процессе существует по крайней мере один поток. Если в процессе существует несколько потоков, они используют одну и ту же память и файловые ресурсы. Потоки не владеют ресурсами, кроме стека и копии регистров. Процессор (cpu) (не процесс, а процессор) работает с потоками. Он выполняет в одни момент времени один поток. Переключение между потоками - это переключение контекста.

// slide with threading

## Threading

Итак, мы поняли, что нам нужно - многопоточность - разделим. Тут показываю код. Всё классно. Но с этим есть очень тяжелые проблемы. Но прежде, можно представить более высокий способ моделирования конкурентного поведения програмы - уже на уровне выполнения самой программы потому что они разделяют те же проблемы.

// slide with classification

Корутины. Чтобы понять зачем они нужны и почему людям стало нехватать использовать разделения на системные потоки, нужно ответить на вопрос..

// slide with question

..Что если я хочу много потоков?

Потому что корутины из-за того что могут быть сделаны на уровне языка программирования гораздо легче чем системные потоки, на которые аллоцируется около 500кб при создании (наиболее экстремальным примером была java, у которого jvm агрессивно аллоцировал до 1мб).

Тут показать в VSCode, что можно легко создать 1_000_000 корутин.

Но это всё круто, но возвращаясь к вопросу - **Зачем мне нужно столько корутин**. Я их никогда в жизни не распараллелю. Но на самом деле моя цель совсем в другом, я хочу смоделировать поведение реальных агентов, о которых я говорил в начале. Я хочу сделать мою программу фундаментально конкурентной. Можно продемонстрировать на самом тупом примере.

Тут показать в VSCode, 


- problems if you mant to allocate lots of threads (Java 8 for example aggresively allocates around 1mb for each thread on creation.)





## Sharing and communicating


Slide with birds sharing:

- also, the only way to communicate is by sharing memory. This can lead to a series of problems..

Slide with cpp and erlang side by side:

..and people don't like to micromanage problems, so later they though up of ways to communicate without sharing memory, but first, problems:

## Problems

### Deadlock

More formal example by Dijkstra

https://en.wikipedia.org/wiki/Dining_philosophers_problem

### Livelock

it's about trying to pass through the door

### Starvation

...

### Race condition

... Undetermined result at the end of execution. Think of real life example. Also present table example from wiki. Hard to debug manually. Example about golang standard library and implemeted race detector.

Real life example - Race condition
Let say that you own a tank of fishes. You feed the fishes every day after you come home from school. But the thing you don't know, is that your sister does the same! When she comes home, she feeds the fishes too. Couple of months and fishes died for overfeeding. Here your unplanned feeding routines were in race condition: you didn't know was the fishes fed or not, so you both did it because you knew that it had to be done.

### Synchronization (one of the solutions to the problem)

...

**Semaphore**

Consider a variable A and a boolean variable S. A is only accessed when S is marked true. Thus, S is a semaphore for A.

Here are the examples of race conditions

**Mutex**

**Queue**

## Actor model and Process calculus

In computer science, the process calculi (or process algebras) are a diverse family of related approaches for formally modelling concurrent systems. Process calculi provide a tool for the high-level description of interactions, communications, and synchronizations between a collection of independent agents or processes. They also provide algebraic laws that allow process descriptions to be manipulated and analyzed, and permit formal reasoning about equivalences between processes.

Two main approaches: **Actor model** and **Process calculus**.

There are many similarities between the two approaches, but the main difference is

- Processes in the process calculi are anonymous, and communicate by sending messages either through named channels (synchronous or asynchronous), or via ambients (which can also be used to model channel-like communications (Cardelli and Gordon 1998)). In contrast, actors in the Actor model possess an identity, and communicate by sending messages to the mailing addresses of other actors (this style of communication can also be used to model channel-like communications—see below).

### Actor model (1973)

https://rocketeer.be/articles/concurrency-in-erlang-scala/

The Actor model was inspired by the laws of physics and depends on them for its fundamental axioms, i.e. physical laws (see Actor model theory); the process calculi were originally inspired by algebra (Milner 1993).

Erlang & Scala

По аналогии с философией объектно-ориентированного программирования, где каждый примитив рассматривается как объект, модель акторов выделяет в качестве универсальной сущности понятие «актора». Актор является вычислительной сущностью, которая в ответ на полученное сообщение может одновременно:

- отправить конечное число сообщений другим акторам;
- создать конечное число новых акторов;
- выбрать поведение, которое будет использоваться при обработке следующего полученного сообщения.

Не предполагается существования определённой последовательности вышеописанных действий и все они могут выполняться параллельно.

Of importance in this model is that all communications are performed asynchronously. This implies that the sender does not wait for a message to be received upon sending it, it immediately continues its execution. There are no guarantees in which order messages will be received by the recipient, but they will eventually be delivered.

A second important property is that all communications happen by means of messages: there is no shared state between actors. If an actor wishes to obtain information about the internal state of another actor, it will have to use messages to request this information. This allows actors to control access to their state, avoiding problems like the lost-update problem. Manipulation of the internal state also happens through messages.

### Erlang

A simple application that uses an actor can be seen below. In this application, an actor is defined which acts as a basic counter. We send 100.000 increment messages to the actor and then request it to print its internal value.

Lines 1 & 2 defines the module and the exported functions. Lines 4 till 7 contain the run function, which starts a counter process and starts sending increment messages. Sending these messages happens in lines 15 till 18, using the message-passing operator (!). As Erlang is a purely functional language, it has no loop structures. Therefore, this has to be expressed using recursion. These extremely deep recursion stacks would lead to stack overflows in Java, yet Erlang is optimized for these usage patterns. The increment message in this example also carries a parameter, to show Erlangs parameter capabilities. The state of the counter is also maintained using recursion: upon receiving an inc message, the counter calls itself with the new value which causes it to receive the next message. If no messages are available yet, the counter will block and wait for the next message.

**Actor scheduling in Erlang**

Erlang uses a preemptive scheduler for the scheduling of processes [7]. When they have executed for a too long period of time (usually measured in the amount of methods invoked or the amount of CPU-cycles used), or when they enter a receive statement with no messages available, the process is halted and placed on a scheduling queue.

This allows for a large number of processes to run, with a certain amount of fairness. Long running computations will not cause other processes to become unresponsive.

### Scala

**Thread-based vs. Event-based actors**

Scala makes the distinction between thread-based and event-based actors.

Thread-based actors are actors which each run in their own JVM thread. They are scheduled by the Java thread scheduler, which uses a preemptive priority-based scheduler. When the actor enters a receive block, the thread is blocked until messages arrive. Thread-based actors make it possible to do long-running computations, or blocking I/O operations inside actors, without hindering the execution of other actors.

There is an important drawback to this method: each thread can be considered as being heavy-weight and uses a certain amount of memory and imposes some scheduling overhead. When large amounts of actors are started, the virtual machine might run out of memory or it might perform suboptimal due to large scheduling overhead.

In situations where this is unacceptable, event-based actors can be used. These actor are not implemented by means of one thread per actor, yet instead they run on the same thread. An actor that waits for a message to be received is not represented by a blocked thread, but by a closure. This closure captures the state of the actor, such that it’s computation can be continued upon receiving a message [10]. The execution of this closure happens on the thread of the sender.

Event-based actors provide a more light-weight alternative, allowing for very large numbers of concurrently running actors. They should however not be used for parallelism: since all actors execute on the same thread, there is no scheduling fairness.


**Safety in Scala concurrency**

Another potential pit-fall in Scala comes from the fact that it mixes actors with object-oriented programming. It is possible to expose the internal state of an actor through publicly available methods for retrieving and modifying this state. When doing so, it is possible to modify an object by directly invoking its methods, that is: without using messages. Doing so means that you no longer enjoy the safety provided by the actor model.

...


## Process Calculus Family

CCS, ACP and **CSP**

channels, man

### CCS (Calculus of Communicating - Robin Milner)

...

### CSP (Communicating sequential processes - Tony Hoare)

Formalized in 1978

![Tony Hoare pic slide]

The same dude that invented Quicksort - used currently in C, C++, Java, Python, where there is no requirement for stable sorting.

*Тут общяя теория про CSP*

**Primitives**

CSP provides two classes of primitives in its process algebra:

Events
    Events represent communications or interactions. They are assumed to be indivisible and instantaneous. They may be atomic names (e.g. on, off), compound names (e.g. valve.open, valve.close), or input/output events (e.g. mouse?xy, screen!bitmap).
Primitive processes
    Primitive processes represent fundamental behaviors: examples include STOP (the process that communicates nothing, also called deadlock), and SKIP (which represents successful termination).

[here is a slide of csp_syntax]

The syntax of CSP defines the “legal” ways in which processes and events may be combined. Let e be an event, and X be a set of events. Then the basic syntax of CSP can be defined as:

**message passing via channels**

Согласно CSP, сначала вводится множество элементарных событий (алфавит), затем из них конструируются процессы, причём из только что описанных процессов можно строить новые. Процессы, протекающие параллельно, обмениваются информацией, используя безбуферный обмен информацией типа «рандеву» между парой (и только парой) процессов посредством специального объекта — канала. При взаимодействии тот участник обмена, который обратился к каналу первым, ожидает готовности партнёра (точки рандеву); при наступлении последней инициируется обмен. Использование общей для нескольких параллельных процессов памяти в CSP не допускается.

It is interesting that with this demands the channel serves two functions - delivering message and synchronizing routines.

Go slightly extends on the idea and adds buffered channels and there is no restriction on the number of clients of the channel.

```go
code example here
```

Эта теория привела к созданию occam в 1983 британской компанией для своих же новых Транспьютеров - элемент построения многопроцессорных систем

**Comparison with the actor model**

- CSP processes are anonymous, while actors have identities.
- CSP uses explicit channels for message passing, whereas actor systems transmit messages to named destination actors. These approaches may be considered duals of each other, in the sense that processes receiving through a single channel effectively have an identity corresponding to that channel, while the name-based coupling between actors may be broken by constructing actors that behave as channels.
- CSP message-passing fundamentally involves a rendezvous between the processes involved in sending and receiving the message, i.e. the sender cannot transmit a message until the receiver is ready to accept it. In contrast, message-passing in actor systems is fundamentally asynchronous, i.e. message transmission and reception do not have to happen at the same time, and senders may transmit messages before receivers are ready to accept them. These approaches may also be considered duals of each other, in the sense that rendezvous-based systems can be used to construct buffered communications that behave as asynchronous messaging systems, while asynchronous systems can be used to construct rendezvous-style communications by using a message/acknowledgement protocol to synchronize senders and receivers.



## Now Go

A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space. It is lightweight, costing little more than the allocation of stack space. And the stacks start small, so they are cheap, and grow by allocating (and freeing) heap storage as required.
Goroutines are multiplexed onto multiple OS threads so if one should block, such as while waiting for I/O, others continue to run. Their design hides many of the complexities of thread creation and management.

Each concurrent piece of code (goroutine) does not need to share memory ( shared variables with locks) with other goroutine like how you would typically approach concurrency in most other languages. Instead, it can share memory via communicating with other goroutines ( this is done by sending the data from one goroutine to another via the go channels). By default, when you pass a message, the two goroutines wait till the message is received on the other end.

The fact that Go supports this natively , covering all the memory management and thread scheduling, makes concurrency easy to implement and efficient in practice from a developer's prospective.

```go
c := make(chan int)

go func() {c <- 42}()

x := <- c
```

### Real life examples

...

## Asynchronous

### But first - what is stack in synchronous code

### Then event loop

Steal from here with slides

https://iximiuz.com/en/posts/explain-event-loop-in-100-lines-of-code/

The most obvious examples are asynchronous I/O, UIs. Because we can't wait. More here on that...

Video or picture of rollercoaster

### In python specifically

https://www.youtube.com/watch?v=tSLDcRkgTsY

What is synchronous code?

Synchronous code:

Is what you're used to!
Runs functions one after another

What is asynchronous code?

Asynchronous code:

Runs multiple functions seemingly in parallel
    In a single process
    Without threads
Requires cooporative, well-behaving functions
    Functions that regularly suspend by awaiting something
Should not use blocking functions!
    No time.sleep()
    No socket.*
    Etc.
    asyncio provides non-blocking alternatives for many of these functions

A note on Python versions

The async and await keywords were introduced in Python 3.5
They are syntactic sugar on top of the asyncio module that was introduced in Python 3.4

Python 3.3 and earlier do not support this

But generator coroutines can do some of the same things

### In javascript

...

http://latentflip.com/loupe/?code=JC5vbignYnV0dG9uJywgJ2NsaWNrJywgZnVuY3Rpb24gb25DbGljaygpIHsKICAgIHNldFRpbWVvdXQoZnVuY3Rpb24gdGltZXIoKSB7CiAgICAgICAgY29uc29sZS5sb2coJ1lvdSBjbGlja2VkIHRoZSBidXR0b24hJyk7ICAgIAogICAgfSwgMjAwMCk7Cn0pOwoKY29uc29sZS5sb2coIkhpISIpOwoKbHNldFRpbWVvdXQoZnVuY3Rpb24gdGltZW91dCgpIHsKICAgIGNvbnNvbGUubG9nKCJDbGljayB0aGUgYnV0dG9uISIpOwp9LCA1MDAwKTsKCmNvbnNvbGUubG9nKCJXZWxjb21lIHRvIGxvdXBlLiIpOw%3D%3D!!!PGJ1dHRvbj5DbGljayBtZSE8L2J1dHRvbj4%3D

### Drawbacks

...



### QA task about race in go
