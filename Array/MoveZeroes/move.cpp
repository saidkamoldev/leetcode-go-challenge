#include <iostream>

int main(){
    int nums[1] = {0};
    int index = 0;
    for(int i = 1; i < 1; i++){
        if(nums[index] != 0 && nums[i] != 0){
            index++;
        }else if(nums[index] == 0 && nums[i] != 0){
            
            nums[index] = nums[i];
            nums[i] = 0;
            index++;
        }
    }
    for(int i : nums){
        std::cout << i;
    }
}