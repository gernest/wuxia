var Wu ={
    tplFuncs:{},
    plugins:{},
    files :fileTree(),
    getTplFuncs:function(){
        return _.map(this.tplFuncs, function(v){return v;});
    },
    prepare:function(plan){
        var that=this;
        if (plan){
            _.each(plan.Dependency, function(el){
               that.plugins[el]= require(el);
            });
            return true;
        }
        return false;
    },
    system: sys,
    getSystem: function(){
        return this.system;
    },
    setCustomPlan: function(plan){
        addCustomPlan(JSON.stringify(plan));
    },
    process: function(strategy,data,out){
        var that=this;
        var file=JSON.parse(data);

        // before
        _.each(strategy.Before,function(s){
            file=require(s).exec(file);
        }) ;

        //exec
        _.each(strategy.Exec,function(s){
            file=require(s).exec(file);
        }) ;

        //after
        _.each(strategy.After,function(s){
            file=require(s).exec(file);
        }) ;
        var v=JSON.stringify(file);
        out.WriteString(v);
    }
};

