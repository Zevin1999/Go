1. Mysql主从复制怎么实现的？ [Binary log/ Relay log]
   主从复制（Replication）是指数据可以从一个MySQL数据库主服务器复制到一个或多个从服务器，从服务器可以复制主服务器中的所有数据库或者特定的数据库，或者特定的表。默认采用异步模式。 
   实现原理： 
   主服务器 binary log dump 线程：将主服务器中的数据更改（增删改）日志写入 Binary log 中；
   从服务器 I/O 线程：负责从主服务器读取binary log，并写入本地的 Relay log；
   从服务器 SQL 线程：负责读取 Relay log，解析出主服务器已经执行的数据更改，并在从服务器中重新执行（Replay），保证主从数据的一致性
2. MySQL的数据如何恢复到任意时间点？
   【已定时的做全量备份，以及备份增量的binlog日志为前提】
   恢复到任意时间点首先将全量备份恢复之后，再此基础上回放增加的binlog直至指定的时间点。
3. 数据库的四种隔离级别?
   1. 未提交读（Read Uncommited）：在一个事务提交之前，它的执行结果对其它事务也是可见的。会导致脏读、不可重复读、幻读 
   2. 提交读（Read Commited）：一个事务只能看见已经提交的事务所作的改变。可避免脏读问题 
   3. 可重复读（Repeatable Read）：可以确保同一个事务在多次读取同样的数据时得到相同的结果。（MySQL的默认隔离级别）。可避免不可重复读 
   4. 可串行化（Serializable）：强制事务串行执行，使之不可能相互冲突，从而解决幻读问题。可能导致大量的超时现象和锁竞争，实际很少使用
4. MVCC如何理解? https://zhuanlan.zhihu.com/p/231947511
   1. 多版本并发控制 (Multi-Version Concurrency Control)，只有在InnoDB引擎下存在。MVCC机制的作用其实就是避免同一个数据在不同事务之间的竞争，提高系统的并发性能。
      它的特点:
      1. 允许多个版本同时存在，并发执行。 
      2. 不依赖锁机制，性能高。 
      3. 只在提交读和可重复读的事务隔离级别下工作。
   2. InnoDB存储引擎中，它的聚簇索引记录中都包含两个必要的隐藏列，分别是： 
      1. trx_id：事务Id，每次一个事务对某条聚簇索引记录进行改动时，都会把该事务的事务id 赋值给trx_id 隐藏列。 
      2. roll_pointer：回滚指针，每次对某条聚簇索引记录进行改动时，都会把旧的版本写入到undo log 中，然后这个隐藏列就相当于一个指针，可以通过它来找到该记录修改前的信息。
   3. ReadView可以理解为数据库中某一个时刻所有未提交事务的快照。ReadView有几个重要的参数： 
      m_ids：表示生成ReadView时，当前系统正在活跃的读写事务的事务Id列表。 
      min_trx_id：表示生成ReadView时，当前系统中活跃的读写事务的最小事务Id。 
      max_trx_id：表示生成ReadView时，当前时间戳InnoDB将在下一次分配的事务id。 
      creator_trx_id：当前事务id。 所以当创建ReadView时，可以知道这个时间点上未提交事务的所有信息。
   4. 事务链每次对记录进行修改时，都会记录一条undo log信息，每一条undo log信息都会有一个roll_pointer属性(INSERT操作没有这个属性，因为之前没有更早的版本)，可以将这些undo日志都连起来，串成一个链表。
   5. 🌟🌟
      如果被访问版本的trx_id属性值小于ReadView的最小事务Id，表示该版本的事务在生成 ReadView 前已经提交，所以该版本可以被当前事务访问。 
      如果被访问版本的trx_id属性值大于ReadView的最大事务Id，表示该版本的事务在生成 ReadView 后才生成，所以该版本不可以被当前事务访问。 
      如果被访问版本的trx_id属性值在m_ids列表最小事务Id和最大事务Id之间，那就需要判断一下 trx_id 属性值是不是包含在 m_ids 列表中，如果包含的话，说明创建 ReadView 时生成该版本的事务还是活跃的，所以该版本不可以访问；如果不包含的话，说明创建 ReadView 时生成该版本的事务已经被提交，该版本可以被访问。
   [MVCC通过对数据进行多版本保存，根据比较版本号来控制数据是否展示，从而达到读取数据时无需加锁就可以实现事务的隔离性。]
5. ACID（参考：https://cloud.tencent.com/developer/article/1888427）
   1. 原子性：单个事务，为一个不可分割的最小工作单元。通过undolog来实现。 
   2. 一致性：数据库总是从一个一致性的状态转换到另外一个一致性的状态。通过binlog、redolog来实现。
   3. 隔离性：通常来说，一个事务所做的修改在最终提交以前，对其他事务是不可见的。通过(读写锁+MVCC)来实现。
   4. 持久性：一旦事务提交，则其所做的修改就会永久保存到数据库中。此时即使系统崩溃，修改的数据也不会丢失。
   MySQL通过原子性，持久性，隔离性最终实现（或者说定义）数据一致性。
逻辑备份日志（binlog）、重做日志（redolog）、回滚日志（undolog）、锁技术 + MVCC就是MySQL实现事务的基础