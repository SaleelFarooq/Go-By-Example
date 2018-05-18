#include <iostream>
#include <string>

using namespace std;

int main()
{ char pos,block;
 int i,j,m,n,res,p,q,dis1,dis2;
  string array1[5]={"ABCDE","FGHIJ","KLMNO","PQRST","UVWXY"};
  cin>>pos>>block;
  
  for(i=0;i<5;i++)
    for(j=0;j<5;j++)
            {if(array1[i][j]==pos)
                    {m=i;
                    n=j;
                    }
            if(array1[i][j]==block)
                    {p=i;
                    q=j;
                    }
            }
if((m==0 && n==0)||(m==0 && n==4)||(m==4 && n==0)||(m==4 && n==4))
    res=3;
else if(m==0 || n==0 || m==4 || n==4)
    res=5;
else
    res=8;
//cout<<res;
dis1=m-p;
if(dis1<0)
    dis1=-1*dis1;
dis2=n-q;
if(dis2<0)
    dis2=-1*dis2;
if((dis1+dis2)<=2)
    res--;
cout<<res;
   return 0;
}
