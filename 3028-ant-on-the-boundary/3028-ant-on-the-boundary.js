/**
 * @param {number[]} nums
 * @return {number}
 */
var returnToBoundaryCount = function(nums) {
    let cou=0;
    let num=0;
    for(let i of nums){
        num+=i;
        if(num==0)cou+=1;
    }
    return cou;
};