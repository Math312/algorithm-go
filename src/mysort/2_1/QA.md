# 2.1 Exercise

## 2.1.1

*Q*: **According to the format of the trajectory shown in Algorithm 2.1, You need to give the procession about how to sort the array "[E , A, S, Y, Q, U, E, S, T, I, O, N]" by selection sort.**

*A*:

index|min| array
---|---|---
1|A|**A**,**E**,S,Y,Q,U,E,S,T,I,O,N
2|E|A,**E**,S,Y,Q,U,E,S,T,I,O,N
3|E|A,E,**E**,Y,Q,U,**S**,S,T,I,O,N
4|I|A,E,E,**I**,Q,U,S,S,T,**Y**,O,N
5|N|A,E,E,I,**N**,U,S,S,T,Y,O,**Q**
6|O|A,E,E,I,N,**O**,S,S,T,Y,**U**,Q
7|Q|A,E,E,I,N,O,**Q**,S,T,Y,U,**S**
8|S|A,E,E,I,N,O,Q,**S**,T,Y,U,S
9|S|A,E,E,I,N,O,Q,S,**S**,Y,U,**T**
10|T|A,E,E,I,N,O,Q,S,S,**T**,U,**Y**
11|U|A,E,E,I,N,O,Q,S,S,T,**U**,Y
12|Y|A,E,E,I,N,O,Q,S,S,T,U,**Y**

## 2.1.2

*Q*: **In selection sorting, how many times can an element be exchanged at most? How many times may be exchanged on average?**

*A*:

When sorting N elements, at most N exchanges are performed, and each element is exchanged once on average.

## 2.1.3

*Q*: **Please provide the array which includes N elements,**

 **构造一个含有个元素的数组，使选择排序（算法2.1）运行过程中a[j]<a[min]（由此min会不断更新）成功的次数最大。**
 
 *A*:
 
 正序数组然后左移一个数字，例如，[1,2,3,4,5]是正序数组，那么答案为[2,3,4,5,1]。按照这种方式创建的数组，每次遍历时最小的数字都在最后一个，因此比较数为0+...+(n-1) = (0+n-1)*n/2
 
 ## 2.1.4
 
*Q*: **按照算法2.2所示轨迹的格式给出插入排序是如何将数组EASYQUESTION排序的**

i|min|0|1|2|3|4|5|6|7|8|9|10|11
---|---|---|---|---|---|---|---|---|---|---|---|---|---
| | |E|A|S|Y|Q|U|E|S|T|I|O|N
0|0|**E**|A|S|Y|Q|U|E|S|T|I|O|N
1|0|**A**|E|S|Y|Q|U|E|S|T|I|O|N
2|2|A|E|**S**|Y|Q|U|E|S|T|I|O|N
3|3|A|E|S|**Y**|Q|U|E|S|T|I|O|N
4|2|A|E|**Q**|S|Y|U|E|S|T|I|O|N
5|4|A|E|Q|S|**U**|Y|E|S|T|I|O|N
6|2|A|E|E|Q|S|**U**|Y|S|T|I|O|N
7|5|A|E|E|Q|S|**S**|U|Y|T|I|O|N
8|6|A|E|E|Q|S|S|**T**|U|Y|I|O|N
9|3|A|E|E|**I**|Q|S|S|T|U|Y|O|N
10|4|A|E|E|I|O|Q|S|S|T|U|Y|N
11|4|A|E|E|I|**N**|O|Q|S|S|T|U|Y

## 2.1.5

*Q*: **构造一个含有个元素的数组，使插入排序（算法2.2）运行过程中内循环（for）的两个判断结果总是假.**

*A*: 同2.1.3，正序数组最小的数字在倒数第一位，例如：[2,3,4,5,1]，比较次数同样是0+...+(n-1)=(0+n-1)*n/2

## 2.1.6

*Q*: **在所有主键都相同的情况下，选择排序和插入排序谁更快？**

*A*: 插入排序更快，在所有键都相等的情况下，插入排序只会比较n次，而选择排序会比较0+...+(n-1)=(0+n-1)*n/2次。

## 2.1.7

*Q*: **对于逆序数组，插入排序和选择排序哪个更快？**

*A*: 一样快，插入排序和选择排序的交换次数相同，都是n次，比较次数也相同，都是0+...+ (n-1)次

## 2.1.8

*Q*: **假设元素只能有3种值，使用插入排序处理这样一个随机数组的运行时间是线性的还是平方级别的？还是介于两者之间？**

*A*: 平方级别的，极少数情况，该数组本身就是有序的，那么是n次比较，否则都是n^2级别的。

## 2.1.9

*Q*: **按照算法2.3所示轨迹的格式给出希尔排序是如何将数组EASYSHELLSORTQUESTION进行排序的。**

*A*: 

## 2.1.10

*Q*: **为什么希尔排序在实现h有序时选用的是插入排序，而不是选择排序？**

*A*: 因为插入排序在部分有序的情况下要比选择排序更快。



