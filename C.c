#include<stdio.h>
int main()
{
    int i,j,s=0,n,a[10000],b[10000],c[10000],d[10000];
    scanf("%d",&n);
    for(i=0;i<n;i++)
    {
       scanf("%d %d",&a[i],&b[i]);
    }
    for(i=0;i<n;i++)
    {
        if(a[i]>=b[i])
        {
            d[i]=a[i];
            c[i]=b[i];
        }
        else if(b[i]>a[i])
        {
            c[i]=a[i];
            d[i]=b[i];
        }
    }
    for(i=0;i<n;i++)
    {
        if(c[i]%2==0)
        {
            for(j=c[i]+1;j<d[i];j=j+2)
            {
               s=s+j;
            }
        }
        if(c[i]%2==1)
        {
            for(j=c[i]+2;j<d[i];j=j+2)
            {
               s=s+j;
            }
        }
        printf("%d\n",s);
        s=0;
    }

    return 0;
}
