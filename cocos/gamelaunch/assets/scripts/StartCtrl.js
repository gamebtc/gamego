
cc.Class({
    extends: cc.Component,

    properties: {
        popupLayer: {
            default: null,
            type: cc.Node
        },
        sceneLayer: {
            default: null,
            type: cc.Node
        }
    },

    start () {
        cc.loader.loadRes("prefabs/loading", cc.Prefab,(err, result)=>{
            if(err){
                return
            }
            let view = cc.instantiate(result);
            this.sceneLayer.addChild(view);
        })

    },

});
