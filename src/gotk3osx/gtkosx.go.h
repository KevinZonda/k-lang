#pragma once

/*
 * This file comes from https://github.com/coyim/gotk3osx/blob/main/gtkosx.h
 */

static GtkosxApplication *
toGtkosxApplication(void *p)
{
    return (GTKOSX_APPLICATION(p));
}

static GtkMenuShell *
toGtkMenuShell(void *p)
{
    return (GTK_MENU_SHELL(p));
}

static GtkMenuItem *
toGtkMenuItem(void *p)
{
    return (GTK_MENU_ITEM(p));
}

static GtkWidget *
toGtkWidget(void *p)
{
    return (GTK_WIDGET(p));
}

static GdkPixbuf *
toGdkPixbuf(void *p)
{
    return (GDK_PIXBUF(p));
}