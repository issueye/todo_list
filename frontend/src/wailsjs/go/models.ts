export namespace db {
	
	export class QryTodoList {
	    group_id: number;
	    filter: string;
	    state: number;
	    date: string;
	
	    static createFrom(source: any = {}) {
	        return new QryTodoList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.group_id = source["group_id"];
	        this.filter = source["filter"];
	        this.state = source["state"];
	        this.date = source["date"];
	    }
	}
	export class TodoGroup {
	    id: number;
	    name: string;
	    description: string;
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    updated_at: any;
	
	    static createFrom(source: any = {}) {
	        return new TodoGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Todo {
	    id: number;
	    title: string;
	    level: string;
	    description: string;
	    completed: boolean;
	    estimated_hours: number;
	    progress: number;
	    // Go type: time
	    due_date: any;
	    group_id: number;
	    group: TodoGroup;
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    updated_at: any;
	
	    static createFrom(source: any = {}) {
	        return new Todo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.level = source["level"];
	        this.description = source["description"];
	        this.completed = source["completed"];
	        this.estimated_hours = source["estimated_hours"];
	        this.progress = source["progress"];
	        this.due_date = this.convertValues(source["due_date"], null);
	        this.group_id = source["group_id"];
	        this.group = this.convertValues(source["group"], TodoGroup);
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class TodoList {
	    all_count: number;
	    uncompleted_count: number;
	    completed_count: number;
	    list: Todo[];
	
	    static createFrom(source: any = {}) {
	        return new TodoList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.all_count = source["all_count"];
	        this.uncompleted_count = source["uncompleted_count"];
	        this.completed_count = source["completed_count"];
	        this.list = this.convertValues(source["list"], Todo);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

