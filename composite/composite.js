const Component = {
    getPrice: function() {},
    print: function(price) {console.log(this.name + ': ' + price)},
};

function Product(name, price) {
    this.name = name;
    this.price = price;
    this.getPrice = function() {
        this.print(this.price);
        return this.price;
    };
};
Product.prototype = Component;

function CompositeProduct(name, children) {
    this.name = name;
    this.children = children;
    this.getPrice = function() {
        let price = 0;
        for (let c of children) {
            price += c.getPrice();
        };
        this.print(price);
        return price;
    };
};
CompositeProduct.prototype = Component;

let turbocharger = new Product('turbocharger', 500);
let intercooler = new Product('intercooler', 200);
let engine_block = new Product('engine_block', 1000);
let engine = new CompositeProduct('engine', [turbocharger, intercooler, engine_block]);

let tire = new Product('tire', 50);
let rim = new Product('rim', 100);
let wheel = new CompositeProduct('wheel', [tire, rim]);

let gearbox = new Product('gearbox', 500);

let body = new Product('body', 1500);

let car = new CompositeProduct('car', [body, engine, gearbox, wheel, wheel, wheel, wheel]);

car.getPrice();
