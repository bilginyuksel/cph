// export declare function test();
// export declare const vari = 100;

function considerCase(callback: ()=>void = () => {console.log("hello world")}) {

}
/**
 * 
 * @param callback 
 */
function considerCase2(callback: (data1, data2) =>void = 
    (data1, data2) => {
        console.log("Hello world");
        const test = () => {
            console.log(data1);
            console.log(data2);
        }
    }) {

}
export async function demo(){}