var searchInsert = function (nums, target) {
  let res = 0;
  if (nums.length % 2 !== 0) {
    res = nums.length - 1;
  } else {
    res = nums.length;
  }
  if (nums.length === 1) {
    if (nums[0] >= target) {
      return 0;
    } else {
      return 1;
    }
  } else if (nums[res / 2] > target) {
    if (nums[0] >= target) {
      return 0;
    }
    for (let i = res / 2; i >= 0; i--) {
      if (nums[i] === target) {
        return i;
      } else if (nums[i] < target && nums[i + 1] > target) {
        return i + 1;
      }
    }
  } else {
    if (nums[nums.length - 1] < target) {
      return nums.length;
    } else if (nums[nums.length - 1] === target) {
      return nums.length - 1;
    }
    for (let i = res / 2; i < res; i++) {
      console.log(i, res, "res");
      if (nums[i] === target) {
        return i;
      } else if (nums[i] < target && nums[i + 1] > target && i !== res) {
        return i + 1;
      }
    }
  }
};
