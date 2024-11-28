#include <iostream>
#include <vector>

int main(){
    std::vector<int> array = {3,2,3,2,2,2,5,5,5,5,5,5,6};
    int count = 0;
    int index = 0;
    for(size_t i = 0; i < array.size(); i++){
        int temp = 0;
        for(size_t j = 0; j < array.size(); j++){
            if(array[i] == array[j]){
                ++temp;
            }
        }
        if(temp > count){
            count = temp;
            index = i;
        }
    }
    std::cout << array[index] << std::endl;
    return 0;
}