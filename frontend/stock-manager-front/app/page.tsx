"use client";

import React, { useState } from 'react';
import { LayoutDashboard, PackagePlus, ClipboardList, ShoppingCart, User } from 'lucide-react';

export default function InventoryApp() {
  const [activeTab, setActiveTab] = useState('shop'); // shop, create, inventory
  const [products, setProducts] = useState([
    { id: 1, name: 'Teclado Mecânico', price: 299.90, stock: 15, category: 'Periféricos' },
    { id: 2, name: 'Mouse Gamer', price: 150.00, stock: 8, category: 'Periféricos' },
  ]);

  // Sidebar Menu Item
  const MenuItem = ({ id, icon: Icon, label }: any) => (
    <button
      onClick={() => setActiveTab(id)}
      className={`w-full flex items-center space-x-3 p-3 rounded-lg transition ${
        activeTab === id ? 'bg-indigo-600 text-white' : 'text-gray-400 hover:bg-gray-800'
      }`}
    >
      <Icon size={20} />
      <span className="font-medium">{label}</span>
    </button>
  );

  return (
    <div className="flex min-h-screen bg-gray-50 text-gray-900">
      {/* --- SIDEBAR --- */}
      <aside className="w-64 bg-gray-900 p-6 flex flex-col">
        <div className="flex items-center space-x-2 mb-10 text-white">
          <div className="bg-indigo-500 p-2 rounded-lg">
            <LayoutDashboard size={24} />
          </div>
          <span className="text-xl font-bold tracking-tight">StockFlow</span>
        </div>

        <nav className="flex-1 space-y-2">
          <MenuItem id="shop" icon={ShoppingCart} label="Loja / Compras" />
          <MenuItem id="inventory" icon={ClipboardList} label="Ver Estoque" />
          <MenuItem id="create" icon={PackagePlus} label="Cadastrar Item" />
        </nav>

        <div className="mt-auto pt-6 border-t border-gray-800 flex items-center space-x-3 text-gray-400">
          <User size={20} />
          <span className="text-sm">Admin v1.0</span>
        </div>
      </aside>

      {/* --- MAIN CONTENT --- */}
      <main className="flex-1 p-10 overflow-y-auto">
        
        {/* VIEW: LOJA / COMPRAS */}
        {activeTab === 'shop' && (
          <section>
            <h2 className="text-3xl font-bold mb-6">Realizar Compras</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {products.map(p => (
                <div key={p.id} className="bg-white p-6 rounded-2xl shadow-sm border border-gray-200">
                  <span className="text-xs font-semibold text-indigo-600 uppercase tracking-wider">{p.category}</span>
                  <h3 className="text-xl font-bold mt-1">{p.name}</h3>
                  <p className="text-2xl font-black text-gray-800 mt-4">R$ {p.price.toFixed(2)}</p>
                  <button className="w-full mt-6 bg-indigo-600 text-white py-3 rounded-xl font-bold hover:bg-indigo-700 transition">
                    Comprar Agora
                  </button>
                </div>
              ))}
            </div>
          </section>
        )}

        {/* VIEW: CADASTRO DE ITENS */}
        {activeTab === 'create' && (
          <section className="max-w-2xl mx-auto bg-white p-8 rounded-2xl shadow-sm border border-gray-200">
            <h2 className="text-2xl font-bold mb-6">Novo Produto</h2>
            <form className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">Nome do Produto</label>
                <input type="text" className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none" placeholder="Ex: Monitor 4K" />
              </div>
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">Preço (R$)</label>
                  <input type="number" className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none" />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">Estoque Inicial</label>
                  <input type="number" className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none" />
                </div>
              </div>
              <button type="button" className="w-full bg-green-600 text-white py-3 rounded-xl font-bold hover:bg-green-700 transition mt-4">
                Salvar no Microserviço (Go)
              </button>
            </form>
          </section>
        )}

        {/* VIEW: VISUALIZAÇÃO DE ESTOQUE */}
        {activeTab === 'inventory' && (
          <section>
            <h2 className="text-3xl font-bold mb-6">Gestão de Inventário</h2>
            <div className="bg-white rounded-2xl shadow-sm border border-gray-200 overflow-hidden">
              <table className="w-full text-left">
                <thead className="bg-gray-50 border-b border-gray-200">
                  <tr>
                    <th className="p-4 font-semibold text-gray-600">ID</th>
                    <th className="p-4 font-semibold text-gray-600">Produto</th>
                    <th className="p-4 font-semibold text-gray-600">Categoria</th>
                    <th className="p-4 font-semibold text-gray-600">Qtd. Atual</th>
                    <th className="p-4 font-semibold text-gray-600">Status</th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-gray-100">
                  {products.map(p => (
                    <tr key={p.id} className="hover:bg-gray-50 transition">
                      <td className="p-4 text-gray-500">#{p.id}</td>
                      <td className="p-4 font-medium">{p.name}</td>
                      <td className="p-4 text-gray-600">{p.category}</td>
                      <td className="p-4">{p.stock} un</td>
                      <td className="p-4">
                        <span className={`px-2 py-1 rounded-full text-xs font-bold ${p.stock > 10 ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-700'}`}>
                          {p.stock > 10 ? 'Saudável' : 'Baixo Estoque'}
                        </span>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </section>
        )}

      </main>
    </div>
  );
}