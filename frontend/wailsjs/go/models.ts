export namespace models {
	
	export class Machine {
	
	
	    static createFrom(source: any = {}) {
	        return new Machine(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

