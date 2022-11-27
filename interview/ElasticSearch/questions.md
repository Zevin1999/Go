1. elasticsearch？ 
   Elasticsearch是通过Lucene的倒排索引技术实现比关系型数据库更快的过滤。它对多条件的过滤支持非常好。
   Elasticsearch 是建立在全文搜索引擎库Lucene基础上的搜索引擎，它隐藏了Lucene的复杂性，取而代之提供了一套简单一致的RESTFUL API，不过掩盖不了它底层也是 Lucene 的事实。Elasticsearch 的倒排索引，其实就是 Lucene 的倒排索引。 
2. elasticsearch 的倒排索引是什么? 
   Elasticsearch使用一种称为倒排索引的结构，它适用于快速的全文搜索。  
   一个倒排索引由文档中所有不重复词的列表构成，对于其中每个词，有一个包含它的文档列表。 
   // ”正向索引”（forward index） "倒排索引"（inverted index） 
   // 全局遍历（效率非常低）——> 排序（二分查找提高遍历效率）——> 跳表数据结构（term dictionary） 
   // 仅使用排序搜索会导致磁盘IO速度过慢，因为数据都放在磁盘中；如果将数据放入内存，又会导致内存爆满。
3. 倒排索引内部结构具体描述
   在term dictionary （word --> documents list）的基础上，新增字典树term index(只存储单词前缀)
   查询流程：通过字典树找到单词块（单词的大概位置），再在单词块里进行二分查找，找到对应的单词，最后找到单词对应的文档列表
   // 为了进一步节省内存，lucene还采用了FST（Finite State Transducers有限状态机）对 Term Index进一步压缩
   // term index在内存中以FST的形式保存，非常节省内存
   // term dictionary在磁盘上以分block的方式保存，block内部利用公共前缀压缩（相比B-Tree更节约磁盘空间）
4. 对查找的Posting List的改进？
   Lucene中，数据根据Segment存储（分片存储），每个Segment最多存65536个文档ID，范围（0～2^16-1)，每个元素占用2个bytes 
   1. 数据压缩（尽可能降低每个数据占用的空间，同时既不让信息失真，又可以还原）
      1. 增量编码（只记录元素与元素之间的增量）
      2. 分割成块（每个块是256个文档ID，增量编码后，每个元素不超过1个字节，也方便求交并集的跳表运算）
      3. 按需分配空间（判断块中最大元素，分配空间）
   2. 快速求交并集
      1. 数组：有序数组，使用跳表（Skip List）。数据需要从磁盘放到内存中处理，内存支撑不住
      2. Bitmap（位图）：0表示角标对应的数字不存在，用1表示存在。节省空间，运算更快
      3. Roaring Bitmaps 算法
   Frame Of Reference（FOR）压缩数据，减少磁盘占用空间
5. 为什么 Elasticsearch/Lucene 检索比mysql快？
   Mysql只有term dictionary这一层，是以 b-tree 排序的方式存储在磁盘上的，检索一个 term 需要若干次随机 IO 的磁盘操作
   Lucene在term dictionary的基础上添加了term index来加速检索，term index 以树的形式缓存在内存中。 
   从 term index 查到对应的 term dictionary 的 block 位置之后，再去磁盘上找 term，大大减少了磁盘的 random access （随机IO）次数。 
   // b+树主要设计目的是减少搜索时访问磁盘的次数
   // Lucene等搜索引擎设计的时候，追求的目标是倒排压缩率&倒排解压速度&倒排Bool运算速度。
   // 取倒排到内存运算的时候，是连续读取，时间开销和倒排的大小有关系，所以并不适合用b+树。
   // Mysql等数据库使用索引的目的是快速定位某一行数据，若使用倒排这种线性化的数据结构存储数据，其查找的时候访问磁盘的次数会远大于使用b+树的数据库。
参考：https://xiaoming.net.cn/2020/11/25/Elasticsearch%20%E5%80%92%E6%8E%92%E7%B4%A2%E5%BC%95/
