function run() {
    const nullData = [{"Ё":0,"А":0,"Б":0,"В":0,"Г":0,"Д":0,"Е":0,"Ж":0,"З":0,"И":0,"Й":0,"К":0,"Л":0,"М":0,"Н":0,"О":0,"П":0,"Р":0,"С":0,"Т":0,"У":0,"Ф":0,"Х":0,"Ц":0,"Ч":0,"Ш":0,"Щ":0,"Ъ":0,"Ы":0,"Ь":0,"Э":0,"Ю":0,"Я":0}]
    webix.ui({
        rows:[
            {/**spacer**/height:10},
            {
                cols:[
                    {/**spacer**/width:20},
                    {
                        id: 'resulted_text',
                        view: "textarea",
                        height: 290,
                        width:660,
                        on:{
                            'onChange': function(id){
                                // webix.message("Text loaded");
                            }
                        },
                    },
                    {/**spacer**/width:10},
                    {
                        rows:[
                            {
                                view:"button",
                                id:"open_button",
                                label:"Open",
                                icon:"fas fa-folder-open",
                                type:"icon",
                                inputWidth:100,
                                on:{
                                    'onItemClick': function(id){
                                        external.invoke('open');
                                    }
                                }
                            },
                            {
                                view:"button",
                                id:"save_button",
                                label:"Save",
                                type:"icon",
                                icon:"fas fa-save",
                                inputWidth:100,
                                on:{
                                    'onItemClick': function(id){
                                        external.invoke('save');
                                    }
                                }
                            },
                            {
                                view:"button",
                                id:"exec_button",
                                label:"Count",
                                type:"icon",
                                icon:"fas fa-table",
                                inputWidth:100,
                                on:{
                                    'onItemClick': function(id){
                                        let text = $$('resulted_text').getValue();
                                        external.invoke('push_table:' + text);
                                    }
                                }
                            },
                            {
                                view:"button",
                                id:"clear_button",
                                label:"Clear",
                                type:"icon",
                                icon:"fas fa-trash-alt",
                                inputWidth:100,
                                on:{
                                    'onItemClick': function(id){
                                        $$('resulted_text').define({value: ""});
                                        $$('resulted_text').refresh();
                                    }
                                }
                            },
                        ],
                    },
                    {/**spacer**/width:20},
                ],
            },
            {/**spacer**/height:10},
            {
                view:"template",
                type:"header",
                template:"Частотный анализ!",
                tip: 'Составить таблицу'
            },
            {
                cols:[
                    {/**spacer**/width:20},
                    {
                        rows:[
                            {/**spacer**/height:10},
                            {
                                autoWidth: false,
                                id: 'datatable_part_1',
                                width: 770,
                                view:"datatable",
                                borders: true,
                                columns:[
                                    { id:"А",    header:"А",   width:70, },
                                    { id:"Б",    header:"Б",   width:70, css: "dark"},
                                    { id:"В",    header:"В",   width:70, },
                                    { id:"Г",    header:"Г",   width:70, css: "dark"},
                                    { id:"Д",    header:"Д",   width:70, },
                                    { id:"Е",    header:"Е",   width:70, css: "dark"},
                                    { id:"Ё",    header:"Ё",   width:70, },
                                    { id:"Ж",    header:"Ж",   width:70, css: "dark"},
                                    { id:"З",    header:"З",   width:70, },
                                    { id:"И",    header:"И",   width:70, css: "dark"},
                                    { id:"Й",    header:"Й",   width:70, },
                                ],
                                data: nullData,
                                scrollX: false,
                                scrollY: false,
                                css: "datatable_style",
                            },
                            {
                                autoWidth: false,
                                id: 'datatable_part_2',
                                width: 770,
                                view:"datatable",
                                columns:[
                                    { id:"К",    header:"К",   width:70, css: "dark"},
                                    { id:"Л",    header:"Л",   width:70},
                                    { id:"М",    header:"М",   width:70, css: "dark"},
                                    { id:"Н",    header:"Н",   width:70},
                                    { id:"О",    header:"О",   width:70, css: "dark"},
                                    { id:"П",    header:"П",   width:70},
                                    { id:"Р",    header:"Р",   width:70, css: "dark"},
                                    { id:"С",    header:"С",   width:70},
                                    { id:"Т",    header:"Т",   width:70, css: "dark"},
                                    { id:"У",    header:"У",   width:70},
                                    { id:"Ф",    header:"Ф",   width:70, css: "dark"},
                                ],
                                data: nullData,
                                scrollX: false,
                                scrollY: false,
                                css: "datatable_style",
                            },
                            {
                                autoWidth: false,
                                id: 'datatable_part_3',
                                width: 770,
                                view:"datatable",
                                columns:[
                                    { id:"Х",   header:"Х",   width:70},
                                    { id:"Ц",   header:"Ц",   width:70, css: "dark"},
                                    { id:"Ч",   header:"Ч",   width:70},
                                    { id:"Ш",   header:"Ш",   width:70, css: "dark"},
                                    { id:"Щ",   header:"Щ",   width:70},
                                    { id:"Ъ",   header:"Ъ",   width:70, css: "dark"},
                                    { id:"Ы",   header:"Ы",   width:70},
                                    { id:"Ь",   header:"Ь",   width:70, css: "dark"},
                                    { id:"Э",   header:"Э",   width:70},
                                    { id:"Ю",   header:"Ю",   width:70, css: "dark"},
                                    { id:"Я",   header:"Я",   width:70},
                                ],
                                data: nullData,
                                scrollX: false,
                                scrollY: false,
                                css: "datatable_style",
                            },
                            {/**spacer**/height:10},
                        ]
                    },
                    {/**spacer**/ width: 20},
                ],
            },
        ],
    });
}

document.addEventListener("DOMContentLoaded", run);