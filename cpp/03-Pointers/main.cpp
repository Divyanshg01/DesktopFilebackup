#include <iostream>
using namespace std;
int main() {
//   int *a_ptr;
//   int *b_ptr{};
//   int *c_ptr{nullptr};
//
// double *a_dptr{};
// double *b_dptr{nullptr};
//   cout << sizeof a_ptr << endl;
//   cout << sizeof b_ptr<< endl;
//   cout << sizeof c_ptr<< endl;
//
//   cout << sizeof a_dptr << endl;
//   cout << sizeof b_dptr<< endl;
    // int num{10};
    // int *ptr{&num};
    // cout<<num<<endl;
    // cout<<sizeof num<<endl;
    // cout<<&num<<endl;
    // cout<<sizeof &num<<endl;
    // cout<< " "<<endl;
    // cout<< " "<<endl;
    // cout<<ptr<<endl;
    // cout<<sizeof ptr<<endl;

   //  int *int_ptr{nullptr};
   //  int_ptr = new int;
   //  cout<<int_ptr<<endl;
   //  cout<<*int_ptr<<endl;
   // *int_ptr = 100;
   //
   //  cout<<int_ptr<<endl;
   //  cout<<*int_ptr<<endl;
   //  delete int_ptr;


    // int *arr_ptr{nullptr};
    // arr_ptr = new int[2];
    //
    //
    //
    // delete [] arr_ptr;
    //


     int x{100};
    int &ref1 = x;
    ref1 = 1000;
    std::cout<<ref1<<endl;
  return 0;
}
