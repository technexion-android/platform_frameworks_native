// Copyright 2024 TechNexion

package inputreader

import (
        "android/soong/android"
        "android/soong/cc"
)

func init() {
    android.RegisterModuleType("inputreader_go", inputreaderDefaultsFactory)
}

func inputreaderDefaultsFactory() (android.Module) {
    module := cc.DefaultsFactory()
    android.AddLoadHook(module, inputreaderDefaults)
    return module
}

func inputreaderDefaults(ctx android.LoadHookContext) {
    type props struct {
        Target struct {
            Android struct {
                Cppflags []string
            }
        }
    }

    p := &props{}
    //SOONG_CONFIG_IMXPLUGIN_TN_MD_TOUCH
    if ctx.Config().VendorConfig("IMXPLUGIN").String("TN_MD_TOUCH") == "true" {
        p.Target.Android.Cppflags = append(p.Target.Android.Cppflags, "-DENABLE_TN_MD_TOUCH")
    }
    ctx.AppendProperties(p)
}
