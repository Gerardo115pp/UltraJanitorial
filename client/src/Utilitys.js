// a function that generates a lorem ipsum string of a given length
export const lorem = (length) => {
        let result = "";
        const words = ["lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscing", "elit", "sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore", "magna", "aliqua."];
        for (let i = 0; i < length; i++) {
            result += result.length > 1 ? words[Math.floor(Math.random() * words.length)] : words[0];
            if (i !== length - 1) {
                result += " ";
            }
        }
        return result;
    };