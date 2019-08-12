from django.urls import path
from rest_framework.urlpatterns import format_suffix_patterns
from snippets import views

from django.conf.urls import include

from snippets.views import SnippetViewSet, UserViewSet

snippet_list = SnippetViewSet.as_view({
    'get': 'list',
    'post': 'create'
})
snippet_detail = SnippetViewSet.as_view({
    'get': 'retrieve',
    'put': 'update',
    'patch': 'partial_update',
    'delete': 'destroy'
})

user_list = UserViewSet.as_view({
    'get': 'list'
})
user_detail = UserViewSet.as_view({
    'get': 'retrieve'
})

# urlpatterns = [
#     path('snippets/', views.SnippetList.as_view(), name='snippet-list'),
#     path('snippets/<int:pk>/', views.SnippetDetail.as_view(), name='snippet-detail'),
#     path('user/', views.UserList.as_view(), name='user-list'),
#     path('user/<int:pk>/', views.UserDetail.as_view(), name='user-detail'),
#     path('', views.api_root),
# ]

urlpatterns = [
    path('snippets/', snippet_list, name='snippet-list'),
    path('snippets/<int:pk>/', snippet_detail, name='snippet-detail'),
    path('user/', user_list, name='user-list'),
    path('user/<int:pk>/', user_list, name='user-detail'),
    path('', views.api_root),
]

urlpatterns += [
    path('api-auth/', include('rest_framework.urls',
                              namespace='rest_framework')),
]

urlpatterns = format_suffix_patterns(urlpatterns)
