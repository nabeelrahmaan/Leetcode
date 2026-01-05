/**
 * @param {number[]} nums
 * @return {number}
 */
var returnToBoundaryCount = function(nums) {
    let count=0;
    let num=0;
    for(i=0;i<nums.length;i++){
        num+=nums[i];
        if(num==0){
            count+=1;
        }
    }
    return count
};