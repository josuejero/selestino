from django.contrib import admin
from django.urls import path
from recipes import views as recipes_views

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', recipes_views.home, name='home'),
    path('search/', recipes_views.search_recipes, name='search_recipes'),
]
