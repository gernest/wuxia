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
};

