// coverage:ignore-file
// ignore_for_file: type=lint
// ignore_for_file: unused_import, prefer_relative_imports, directives_ordering

// GENERATED CODE - DO NOT MODIFY BY HAND

// **************************************************************************
// AppGenerator
// **************************************************************************

// ignore_for_file: no_leading_underscores_for_library_prefixes
import 'package:widgetbook/widgetbook.dart' as _i1;
import 'package:widgetbook_workspace/examples/cool_button.dart' as _i2;
import 'package:widgetbook_workspace/presentation/page/rooms_list_page.dart'
    as _i3;

final directories = <_i1.WidgetbookNode>[
  _i1.WidgetbookFolder(
    name: 'examples',
    children: [
      _i1.WidgetbookLeafComponent(
        name: 'CoolButton',
        useCase: _i1.WidgetbookUseCase(
          name: 'Deafult',
          builder: _i2.buildCoolButtonUseCase,
        ),
      )
    ],
  ),
  _i1.WidgetbookFolder(
    name: 'presentation',
    children: [
      _i1.WidgetbookFolder(
        name: 'page',
        children: [
          _i1.WidgetbookLeafComponent(
            name: 'RoomsListPage',
            useCase: _i1.WidgetbookUseCase(
              name: 'Deafult',
              builder: _i3.buildCoolButtonUseCase,
            ),
          )
        ],
      )
    ],
  ),
];
