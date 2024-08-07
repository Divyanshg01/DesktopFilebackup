#include <iostream>
using namespace std;
int main(){
    int x = 0;
    if (x == 5){
        cout<<"x is 5"<<endl;
    }
    else{
        cout<<"x is not 5 OK"<<endl;
    }
    for(int i =0 ;i<10;i++){
        x = i + x*i;
    }
    return 0;
}
