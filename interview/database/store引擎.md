1. 存储引擎 
   1. ISAM(索引顺序访问方法)：ISAM实现的数据结构是多叉平衡树。但是不支持事务处理(如果数据库崩溃，数据文件无法恢复)，不支持外键，不支持索引。允许没有任何索引和主键的表存在 
   2. MYISAM：支持索引，但不支持事务处理和外键。使用表级锁，来优化多个并发的读写操作。但是需要经常运行optimize 表名恢复更新操作浪费的空间（因为数据库在删除等更新操作后，会在库中留下碎片，影响访问性能，所以必须采用optimize语句进行碎片整理） 
   3. HEAP：Heap将数据存储在内存中，没有磁盘的IO操作，插入、更新查询速度很快。但是保存的数据不稳定(重启或者计算机关机就会消失)，所以这种存储引擎中的表的生命周期很短，一般只使用一次。 
   4. InnoDB：要比ISAM和MyISAM引擎慢很多， InnoDB为MySQL表提供了ACID事务支持、系统崩溃修复能力和多版本并发控制(MVCC)的行级锁，该引擎还提供了行级锁和外键约束，所以InnoDB是事务型数据库的首选引擎。采用B+树实现，索引与数据存储在同一文件中。
2. InnoDB和MyISAM的区别
   1. 存储结构
      1. InnoDB: 共享表空间，所有的数据都存储在一个单独的表空间中，一个表可以跨多个文件存在。
      2. MyISAM: 每个表存储在三个分离的文件中，每一个文件的名字均以表的名字开始，扩展名指出文件类型：.frm文件存储表定义；·MYD (MYData)文件存储表的数据；.MYI (MYIndex)文件存储表的索引
   2. 存储空间
      1. InnoDB存储引擎为在主内存建立其专用的缓冲池来缓存数据和索引，所以需要更多的内存和存储。
      2. MyISAM可被压缩(压缩的表是只读的)，存储空间较小。支持三种不同的存储格式：静态表(默认，但是注意数据末尾不能有空格，会被去掉)、动态表、压缩表。
   3. 索引和数据
      1. InnoDB: 存储索引的是B+树，叶子节点的data域保存了完整的数据记录，以主键作为键，其他列的值作为data域。所以数据本身就是索引文件。InnoDB要求表必须有主键，否则会自动选择某一列作为主键；InnoDB的辅助索引data域存储的是主键的值，所以当以辅助索引查找时，会先根据辅助索引找到主键，再根据主键索引找到实际的数据。 
      2. MYISAM：存储索引的是B+树，其叶子节点以主键作为键，data域存储的相应行记录的地址，因此索引和数据文件是分开的。在MyISAM中，主索引和辅助索引（Secondary key）在结构上没有任何区别，只是主索引要求key是唯一的，而辅助索引的key可以重复。与InnoDB不同的是MYISAM辅助索引存储的data域存储的是对应行的地址而不是主键。
   4. 锁
      1. InnoDB: 支持行级锁。如果执行某个SQL语句时，不确定要扫描的范围那么就会锁定全表。 
      2. MYISAM：支持表级锁。
   5. 事务
      1. InnoDB: 事务安全型，支持外键 
      2. MYISAM：非事务安全型，不支持事务，不支持外键
   6. 全文索引
      1. InnoDB: 支持全文索引。 
      2. MYISAM：不支持全文索引。
   7. 是否保存行数 
      1. InnoDB中不保存表的具体行数，执行SELECT COUNT(*) FROM TABLE时，InnoDB要扫描一遍整个表来计算有多少行。 
      2. MyISAM中存储了表的行数，于是SELECT COUNT(*) FROM TABLE时只需要直接读取已经保存好的值而不需要进行全表扫描
   8. 应用场景
      1. InnoDB: 用于事务处理应用程序，如果应用中需要执行大量的插入删除操作就要用InnoDB,提高并发操作的性能。
      2. MYISAM：管理非事务表，提供高速存储和查询操作。如果要进行大量select操作时适用。
参考链接：https://www.cnblogs.com/winterfells/p/9432217.html