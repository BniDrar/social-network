const throttle = (func, wait) => {
    let called = false;
    return (...args) => {
        if (!called) {
            func.apply(this, args);
            called = true
            setTimeout(() => {
                called = false
            }, wait)
        }
    }
};

export { throttle };
