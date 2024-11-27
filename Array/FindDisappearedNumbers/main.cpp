#include <iostream>
#include <vector>
#include <algorithm>
int main(){
    std::vector<int> nums = {4,3,2,7,8,2,3,1};
    int temp = 1;
    std::vector<int> answer;
    std::sort(nums.begin(), nums.end());
    for(int i = 0; i < nums.size(); i++){
        if(i > 0 && nums[i] == nums[i - 1]){
            continue;
        }
        if(nums[i] != temp){
            answer.push_back(temp);
        }
        ++temp;
    }
    for(int i : answer){
        std::cout << i << " ";
    }
    return 0;

}